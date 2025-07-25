package generator

import (
	"context"
	"os"
	"path/filepath"

	"github.com/openshift/assisted-service/internal/common"
	eventsapi "github.com/openshift/assisted-service/internal/events/api"
	"github.com/openshift/assisted-service/internal/ignition"
	"github.com/openshift/assisted-service/internal/installercache"
	manifestsapi "github.com/openshift/assisted-service/internal/manifests/api"
	"github.com/openshift/assisted-service/internal/provider/registry"
	logutil "github.com/openshift/assisted-service/pkg/log"
	"github.com/openshift/assisted-service/pkg/s3wrapper"
	"github.com/sirupsen/logrus"
)

//go:generate mockgen --build_flags=--mod=mod -package generator -destination mock_install_config.go . InstallConfigGenerator
type InstallConfigGenerator interface {
	GenerateInstallConfig(ctx context.Context, cluster common.Cluster, cfg []byte, releaseImage, installerReleaseImageOverride string, forceInsecurePolicyJson bool) error
}

type Config struct {
	WorkDir            string `envconfig:"WORK_DIR" default:"/data/"`
	ServiceCACertPath  string `envconfig:"SERVICE_CA_CERT_PATH" default:""`
	ReleaseImageMirror string `envconfig:"OPENSHIFT_INSTALL_RELEASE_IMAGE_MIRROR" default:""`
	DummyIgnition      bool   `envconfig:"DUMMY_IGNITION"`
	InstallInvoker     string `envconfig:"INSTALL_INVOKER" default:"assisted-installer"`

	// Directory containing pre-generated TLS certs/keys for the ephemeral installer
	ClusterTLSCertOverrideDir string `envconfig:"EPHEMERAL_INSTALLER_CLUSTER_TLS_CERTS_OVERRIDE_DIR" default:""`
}

type installGenerator struct {
	Config
	log              logrus.FieldLogger
	s3Client         s3wrapper.API
	workDir          string
	providerRegistry registry.ProviderRegistry
	manifestApi      manifestsapi.ManifestsAPI
	eventsHandler    eventsapi.Handler
	installerCache   *installercache.Installers
}

// GetWorkingDirectory determines the working directory path for the generator.
func (c *Config) GetWorkingDirectory() string {
	return filepath.Join(c.WorkDir, "install-config-generate")
}

func New(log logrus.FieldLogger, s3Client s3wrapper.API, cfg Config,
	providerRegistry registry.ProviderRegistry, manifestApi manifestsapi.ManifestsAPI, eventsHandler eventsapi.Handler, installerCache *installercache.Installers) *installGenerator {
	return &installGenerator{
		Config:           cfg,
		log:              log,
		s3Client:         s3Client,
		workDir:          cfg.GetWorkingDirectory(),
		providerRegistry: providerRegistry,
		manifestApi:      manifestApi,
		eventsHandler:    eventsHandler,
		installerCache:   installerCache,
	}
}

// GenerateInstallConfig creates install config and ignition files
func (k *installGenerator) GenerateInstallConfig(ctx context.Context, cluster common.Cluster, cfg []byte, releaseImage, installerReleaseImageOverride string, forceInsecurePolicyJson bool) error {
	log := logutil.FromContext(ctx, k.log)
	err := os.MkdirAll(k.workDir, 0o755)
	if err != nil {
		return err
	}
	clusterWorkDir, err := os.MkdirTemp(k.workDir, cluster.ID.String()+".")
	if err != nil {
		return err
	}
	defer func() {
		if removeError := os.RemoveAll(clusterWorkDir); removeError != nil {
			log.WithError(removeError).Error("Failed to clean up generated ignition directory")
		}
	}()

	// runs openshift-install to generate ignition files, then modifies them as necessary
	var generator ignition.Generator
	if k.Config.DummyIgnition {
		generator = ignition.NewDummyGenerator(clusterWorkDir, &cluster, k.s3Client, log)
	} else {
		generator = ignition.NewGenerator(clusterWorkDir, &cluster, releaseImage, k.Config.ReleaseImageMirror,
			k.Config.ServiceCACertPath, k.Config.InstallInvoker, k.s3Client, log, k.providerRegistry, installerReleaseImageOverride, k.Config.ClusterTLSCertOverrideDir, k.manifestApi, k.eventsHandler, k.installerCache)
	}
	err = generator.Generate(ctx, cfg, forceInsecurePolicyJson)
	if err != nil {
		return err
	}

	// upload files to S3
	err = generator.UploadToS3(ctx)
	if err != nil {
		return err
	}

	return nil
}
