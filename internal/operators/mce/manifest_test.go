package mce

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/openshift/assisted-service/internal/common"
	"github.com/openshift/assisted-service/models"
	"sigs.k8s.io/yaml"
)

var _ = Describe("MCE manifest generation", func() {
	operator := NewMceOperator(common.GetTestLog())
	var cluster *common.Cluster

	getCluster := func(openshiftVersion string) *common.Cluster {
		cluster := common.Cluster{Cluster: models.Cluster{
			OpenshiftVersion: openshiftVersion,
		}}
		return &cluster
	}

	Context("MCE Manifest", func() {
		It("Get MCE channel", func() {
			Expect(getSubscriptionChannel("4.14")).To(Equal("stable-2.4"))
			Expect(getSubscriptionChannel("4.13")).To(Equal("stable-2.3"))
			Expect(getSubscriptionChannel("4.10")).To(Equal("stable-2.0"))
			Expect(getSubscriptionChannel("4.12.0-0.nightly-2022-10-25-210451")).To(Equal("stable-2.2"))
			Expect(getSubscriptionChannel("4.11.0-ec.3")).To(Equal("stable-2.1"))
		})

		It("Check YAMLs of MCE", func() {
			cluster = getCluster("4.10.17")
			openshiftManifests, manifest, err := operator.GenerateManifests(cluster)

			Expect(err).ShouldNot(HaveOccurred())
			Expect(openshiftManifests).To(HaveLen(3))
			Expect(openshiftManifests["50_openshift-mce_ns.yaml"]).NotTo(HaveLen(0))
			Expect(openshiftManifests["50_openshift-mce_operator_subscription.yaml"]).NotTo(HaveLen(0))
			Expect(openshiftManifests["50_openshift-mce_operator_group.yaml"]).NotTo(HaveLen(0))

			for _, manifest := range openshiftManifests {
				_, err = yaml.YAMLToJSON(manifest)
				Expect(err).ShouldNot(HaveOccurred())
			}

			_, err = yaml.YAMLToJSON(manifest)
			Expect(err).ShouldNot(HaveOccurred(), "yamltojson err: %v", err)
		})
	})
})
