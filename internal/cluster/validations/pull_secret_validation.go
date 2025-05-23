package validations

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/openshift/assisted-service/pkg/auth"
	"github.com/openshift/assisted-service/pkg/mirrorregistries"
	"github.com/openshift/assisted-service/pkg/ocm"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"golang.org/x/exp/maps"
)

const (
	dockerHubRegistry   = "docker.io"
	dockerHubLegacyAuth = "https://index.docker.io/v1/"
	stageRegistry       = "registry.stage.redhat.io"
)

// PullSecretValidator is used run validations on a provided pull secret
// it verifies the format of the pull secret and access to required image registries
//
//go:generate mockgen -source=pull_secret_validation.go -package=validations -destination=mock_pull_secret_validation.go
type PullSecretValidator interface {
	ValidatePullSecret(additionalPublicRegistries []string, secret string, username string, releaseImageURL string) error
}

func ParsePublicRegistries(publicRegistries map[string]bool, publicRegistriesLiteral string) {
	if publicRegistriesLiteral == "" {
		return
	}

	for _, registry := range strings.Split(publicRegistriesLiteral, ",") {
		publicRegistries[registry] = true
	}
}

func ParseMirrorRegistries(log logrus.FieldLogger, publicRegistries map[string]bool, mirrorRegistries []mirrorregistries.RegistriesConf) {
	for _, mirrorRegistry := range mirrorRegistries {
		registry := ParseBaseRegistry(mirrorRegistry.Location)
		publicRegistries[registry] = true
	}
}

type registryPullSecretValidator struct {
	publicRegistries   map[string]bool
	registriesWithAuth map[string]bool
	authHandler        auth.Authenticator
}

type imagePullSecret struct {
	Auths map[string]map[string]interface{} `json:"auths"`
}

type PullSecretCreds struct {
	Username string
	Password string
	Registry string
	AuthRaw  string
	Email    string
}

// PullSecretError distinguishes secret validation errors produced by this package from other types of errors
type PullSecretError struct {
	Msg   string
	Cause error
}

func (e *PullSecretError) Error() string {
	return e.Msg
}

func (e *PullSecretError) Unwrap() error {
	return e.Cause
}

// ParsePullSecret validates the format of a pull secret and converts the secret string into individual credential entries
func ParsePullSecret(secret string) (map[string]PullSecretCreds, error) {
	result := make(map[string]PullSecretCreds)
	var s imagePullSecret

	err := json.Unmarshal([]byte(strings.TrimSpace(secret)), &s)
	if err != nil {
		return nil, &PullSecretError{Msg: "pull secret must be a well-formed JSON", Cause: err}
	}
	if len(s.Auths) == 0 {
		return nil, &PullSecretError{Msg: "pull secret must contain 'auths' JSON-object field"}
	}

	for d, a := range s.Auths {
		pullSecretCreds, err := parseAuthConfig(d, a)
		if err != nil {
			return nil, err
		}

		result[d] = pullSecretCreds
	}

	return result, nil
}

func AddRHRegPullSecret(secret, rhCred string) (string, error) {
	if rhCred == "" {
		return "", errors.Errorf("invalid pull secret")
	}
	var s imagePullSecret
	err := json.Unmarshal([]byte(strings.TrimSpace(secret)), &s)
	if err != nil {
		return secret, errors.Errorf("invalid pull secret: %v", err)
	}
	s.Auths[stageRegistry] = make(map[string]interface{})
	s.Auths[stageRegistry]["auth"] = base64.StdEncoding.EncodeToString([]byte(rhCred))
	ps, err := json.Marshal(s)
	if err != nil {
		return secret, err
	}
	return string(ps), nil
}

func NewPullSecretValidator(publicRegistries map[string]bool, authHandler auth.Authenticator, images ...string) (PullSecretValidator, error) {
	registriesWithAuth := map[string]bool{}
	for _, image := range images {
		registryWithAuth, err := getRegistryAuthStatus(publicRegistries, image)
		if err != nil {
			return nil, err
		}

		if registryWithAuth != nil {
			registriesWithAuth[*registryWithAuth] = true
		}
	}

	return &registryPullSecretValidator{
		publicRegistries:   publicRegistries,
		registriesWithAuth: registriesWithAuth,
		authHandler:        authHandler,
	}, nil
}

func parseAuthConfig(registry string, authConfig map[string]interface{}) (PullSecretCreds, error) {
	email, _ := authConfig["email"].(string)

	ret := PullSecretCreds{
		Registry: registry,
		Email:    email,
	}

	auth, authPresent := authConfig["auth"]
	if !authPresent {
		return ret, nil
	}

	authRaw, ok := auth.(string)
	if !ok {
		return ret, &PullSecretError{Msg: fmt.Sprintf("invalid pull secret: 'auth' field of %q is %v but should be a string", registry, authConfig["auth"])}
	}

	if len(authRaw) == 0 {
		return ret, nil
	}

	ret.AuthRaw = authRaw

	data, err := base64.StdEncoding.DecodeString(authRaw)
	if err != nil {
		return ret, &PullSecretError{Msg: fmt.Sprintf("invalid pull secret: 'auth' field of %q is not base64-encoded", registry)}
	}

	res := strings.SplitN(string(data), ":", 2)
	if len(res) != 2 {
		return ret, &PullSecretError{Msg: fmt.Sprintf("invalid pull secret: 'auth' for %s is not in 'user:password' format", registry)}
	}

	ret.Username = res[0]
	ret.Password = res[1]

	return ret, nil
}

// ValidatePullSecret validates that a pull secret is well-formed and contains all of the required data
func (v *registryPullSecretValidator) ValidatePullSecret(additionalPublicRegistries []string, secret string, username string, releaseImageURL string) error {
	creds, err := ParsePullSecret(secret)
	if err != nil {
		return err
	}

	// only check for cloud creds if we're authenticating against Red Hat SSO
	if v.authHandler.AuthType() == auth.TypeRHSSO {
		r, ok := creds["cloud.openshift.com"]
		if !ok {
			return &PullSecretError{Msg: "pull secret must contain auth for \"cloud.openshift.com\""}
		}

		var user interface{}
		user, err = v.authHandler.AuthAgentAuth(r.AuthRaw)
		if err != nil {
			return &PullSecretError{Msg: "failed to authenticate the pull secret token"}
		}

		if (user.(*ocm.AuthPayload)).Username != username {
			return &PullSecretError{Msg: "pull secret token does not match current user"}
		}
	}

	ignorableRegistries := maps.Clone(v.publicRegistries)
	for _, registry := range additionalPublicRegistries {
		ignorableRegistries[registry] = true
	}

	registries := maps.Keys(v.registriesWithAuth)
	registries = append(registries, releaseImageURL)
	return verifyNoMissingAuth(ignorableRegistries, creds, registries...)
}

// verifyNoMissingAuth verifies that all registries are accounted for in either the ignore list or contain
// an authentication entry in the pull secret
func verifyNoMissingAuth(ignorableRegistries map[string]bool, credentials map[string]PullSecretCreds, registries ...string) error {
	// Combine ignorable registries and the registries listed in the credentials since they effectively
	// are the same check: if the registry exists in the ignore list or in the credentials list, then the
	// registry does not require auth
	for key := range credentials {
		ignorableRegistries[key] = true
	}
	for _, registry := range registries {
		authStillRequired, err := getRegistryAuthStatus(ignorableRegistries, registry)
		if err != nil {
			return err
		}

		if authStillRequired != nil {
			return &PullSecretError{Msg: fmt.Sprintf("pull secret must contain auth for %q", registry)}
		}
	}
	return nil
}

// getRegistryAuthStatus takes a release image reference and a set of ignorable registries,
// and returns the image's registry if it requires authentication
func getRegistryAuthStatus(ignorableRegistries map[string]bool, image string) (*string, error) {
	if image == "" {
		return nil, nil
	}

	registry, err := ParseFullRegistry(image)
	if err != nil {
		return nil, errors.Wrapf(err, "error occurred while trying to parse the registry out of '%s'", image)
	}

	if findFullRegistryInSet(registry, ignorableRegistries) {
		return nil, nil
	}
	return &registry, nil
}

// findFullRegistryInSet takes a full image registry (base domain + path) and a set of registries
// and checks if the image registry exists in the set by first checking if the full registry name
// is in the set. If it isn't, it truncates the path of the registry and checks the remainder
// against the set. It continues to do this all the way to the base domain of the registry.
// E.g. quay.io/test/image first checks for quay.io/test/image, then quay.io/test, and finally quay.io
// If any of these exist in the set, this function will return true
func findFullRegistryInSet(registry string, set map[string]bool) bool {
	if isRegistryInSet(registry, set) {
		return true
	}
	for slash := strings.LastIndex(registry, "/"); slash > 0; slash = strings.LastIndex(registry, "/") {
		registry = registry[:slash]
		if isRegistryInSet(registry, set) {
			return true
		}
	}

	return false
}

func isRegistryInSet(registry string, set map[string]bool) bool {
	// Both "docker.io" and "https://index.docker.io/v1/" are acceptable for DockerHub login
	if registry == dockerHubRegistry {
		if _, ok := set[dockerHubLegacyAuth]; ok {
			return true
		}
	}

	// We add auth for stage registry automatically
	return registry == stageRegistry || set[registry]
}
