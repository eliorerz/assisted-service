package subsystem

import (
	"context"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/openshift/assisted-service/cmd/agentbasedinstaller"
	"github.com/openshift/assisted-service/internal/common"
	"github.com/openshift/assisted-service/internal/network"
	"github.com/openshift/assisted-service/subsystem/utils_test"
)

// Note: utils_test.TestContext.UserBMClient is used because subsystems defaults to use "rhsso" as AUTH_TYPE.
// The ephermeral installer environment will use the "none" AUTH_TYPE at the start, and
// a pre-generated infra-env-id will be used when creating the infra-env.
// A new authentication scheme suited to the agent-based installer will be implemented
// in the future and utils_test.TestContext.UserBMClient should be replaced at that time.
var _ = Describe("RegisterClusterAndInfraEnv", func() {
	ctx := context.Background()

	It("good flow", func() {
		modelCluster, registerClusterErr := agentbasedinstaller.RegisterCluster(ctx, log, utils_test.TestContext.UserBMClient, pullSecret,
			"../docs/hive-integration/crds/clusterDeployment.yaml",
			"../docs/hive-integration/crds/agentClusterInstall.yaml",
			"../docs/hive-integration/crds/clusterImageSet.yaml", "", "", false)
		Expect(registerClusterErr).NotTo(HaveOccurred())
		Expect(network.GetApiVipById(&common.Cluster{Cluster: *modelCluster}, 0)).To(Equal("1.2.3.8"))
		Expect(network.GetIngressVipById(&common.Cluster{Cluster: *modelCluster}, 0)).To(Equal("1.2.3.9"))
		Expect(modelCluster.OpenshiftVersion).To(ContainSubstring("4.15.0"))
		Expect(modelCluster.CPUArchitecture).To(Equal("x86_64"))
		Expect(modelCluster.Name).To(Equal("test-cluster"))

		modelInfraEnv, registerInfraEnvErr := agentbasedinstaller.RegisterInfraEnv(ctx, log, utils_test.TestContext.UserBMClient, pullSecret,
			modelCluster, "../docs/hive-integration/crds/infraEnv.yaml",
			"../docs/hive-integration/crds/nmstate.yaml", "full-iso", "")

		Expect(registerInfraEnvErr).NotTo(HaveOccurred())
		Expect(*modelInfraEnv.Name).To(Equal("myinfraenv"))
		Expect(modelInfraEnv.ClusterID).To(Equal(*modelCluster.ID))
		Expect(len(modelInfraEnv.StaticNetworkConfig)).ToNot(BeZero())
	})

	It("InstallConfig override good flow", func() {
		modelCluster, registerClusterErr := agentbasedinstaller.RegisterCluster(ctx, log, utils_test.TestContext.UserBMClient, pullSecret,
			"../docs/hive-integration/crds/clusterDeployment.yaml",
			"../docs/hive-integration/crds/agentClusterInstall-with-installconfig-overrides.yaml",
			"../docs/hive-integration/crds/clusterImageSet.yaml", "", "", false)
		Expect(registerClusterErr).NotTo(HaveOccurred())
		Expect(network.GetApiVipById(&common.Cluster{Cluster: *modelCluster}, 0)).To(Equal("1.2.3.8"))
		Expect(network.GetIngressVipById(&common.Cluster{Cluster: *modelCluster}, 0)).To(Equal("1.2.3.9"))
		Expect(modelCluster.OpenshiftVersion).To(ContainSubstring("4.15.0"))
		Expect(modelCluster.CPUArchitecture).To(Equal("x86_64"))
		Expect(modelCluster.InstallConfigOverrides).To(Equal(`{"fips": true}`))
		Expect(modelCluster.Name).To(Equal("test-cluster"))

		modelInfraEnv, registerInfraEnvErr := agentbasedinstaller.RegisterInfraEnv(ctx, log, utils_test.TestContext.UserBMClient, pullSecret,
			modelCluster, "../docs/hive-integration/crds/infraEnv.yaml",
			"../docs/hive-integration/crds/nmstate.yaml", "full-iso", "")

		Expect(registerInfraEnvErr).NotTo(HaveOccurred())
		Expect(*modelInfraEnv.Name).To(Equal("myinfraenv"))
		Expect(modelInfraEnv.ClusterID).To(Equal(*modelCluster.ID))
		Expect(len(modelInfraEnv.StaticNetworkConfig)).ToNot(BeZero())
	})

	It("missing one of the ZTP manifests", func() {
		modelCluster, registerClusterErr := agentbasedinstaller.RegisterCluster(ctx, log, utils_test.TestContext.UserBMClient, pullSecret,
			"file-does-not-exist",
			"../docs/hive-integration/crds/agentClusterInstall.yaml",
			"../docs/hive-integration/crds/clusterImageSet.yaml", "", "", false)
		Expect(registerClusterErr).To(HaveOccurred())
		Expect(modelCluster).To(BeNil())
	})
})
