package isoeditor

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/cavaliercoder/go-cpio"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/openshift/assisted-service/internal/constants"
	"github.com/openshift/assisted-service/pkg/staticnetworkconfig"
)

func TestIsoEditor(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "IsoEditor")
}

var _ = Describe("RamdiskImageArchive", func() {
	It("adds a new archive correctly - ocp versions less than MinimalVersionForNmstatectl", func() {

		clusterProxyInfo := ClusterProxyInfo{
			HTTPProxy:  "http://10.10.1.1:3128",
			HTTPSProxy: "https://10.10.1.1:3128",
			NoProxy:    "quay.io",
		}

		staticnetworkConfigOutput := []staticnetworkconfig.StaticNetworkConfigData{
			{
				FilePath:     "1.nmconnection",
				FileContents: "1.nmconnection contents",
			},
			{
				FilePath:     "2.nmconnection",
				FileContents: "2.nmconnection contents",
			},
		}

		archive, err := RamdiskImageArchive(staticnetworkConfigOutput, &clusterProxyInfo, constants.PreNetworkConfigScript, constants.MinimalISONetworkConfigService)
		Expect(err).ToNot(HaveOccurred())

		By("checking that the files are present in the archive")
		gzipReader, err := gzip.NewReader(bytes.NewReader(archive))
		Expect(err).ToNot(HaveOccurred())

		var serviceLink, rootfsServiceConfigContent string
		r := cpio.NewReader(gzipReader)
		for {
			hdr, err := r.Next()
			if err == io.EOF {
				break
			}
			Expect(err).ToNot(HaveOccurred())
			switch hdr.Name {
			case "/etc/1.nmconnection":
				configBytes, err := io.ReadAll(r)
				Expect(err).ToNot(HaveOccurred())
				Expect(string(configBytes)).To(Equal("1.nmconnection contents"))
			case "/etc/2.nmconnection":
				configBytes, err := io.ReadAll(r)
				Expect(err).ToNot(HaveOccurred())
				Expect(string(configBytes)).To(Equal("2.nmconnection contents"))
			case "/usr/local/bin/pre-network-manager-config.sh":
				scriptBytes, err := io.ReadAll(r)
				Expect(string(scriptBytes)).To(Equal(constants.PreNetworkConfigScript))
				Expect(err).ToNot(HaveOccurred())
			case "/etc/systemd/system/pre-network-manager-config.service":
				serviceBytes, err := io.ReadAll(r)
				Expect(string(serviceBytes)).To(Equal(constants.MinimalISONetworkConfigService))
				Expect(err).ToNot(HaveOccurred())
			case "/etc/systemd/system/initrd.target.wants/pre-network-manager-config.service":
				Expect(hdr.Mode & cpio.ModeSymlink).ToNot(BeZero())
				serviceLink = hdr.Linkname
			case "/etc/systemd/system/coreos-livepxe-rootfs.service.d/10-proxy.conf":
				rootfsServiceConfigBytes, err := io.ReadAll(r)
				Expect(err).ToNot(HaveOccurred())
				rootfsServiceConfigContent = string(rootfsServiceConfigBytes)
			}
		}

		Expect(serviceLink).To(Equal("/etc/systemd/system/pre-network-manager-config.service"))

		rootfsServiceConfig := fmt.Sprintf("[Service]\n"+
			"Environment=http_proxy=%s\nEnvironment=https_proxy=%s\nEnvironment=no_proxy=%s\n"+
			"Environment=HTTP_PROXY=%s\nEnvironment=HTTPS_PROXY=%s\nEnvironment=NO_PROXY=%s",
			clusterProxyInfo.HTTPProxy, clusterProxyInfo.HTTPSProxy, clusterProxyInfo.NoProxy,
			clusterProxyInfo.HTTPProxy, clusterProxyInfo.HTTPSProxy, clusterProxyInfo.NoProxy)
		Expect(rootfsServiceConfigContent).To(Equal(rootfsServiceConfig))
	})
	It("returns nothing when given nothing - ocp versions less than MinimalVersionForNmstatectl", func() {
		archive, err := RamdiskImageArchive(
			[]staticnetworkconfig.StaticNetworkConfigData{},
			&ClusterProxyInfo{},
			constants.PreNetworkConfigScript, constants.MinimalISONetworkConfigService)
		Expect(err).ToNot(HaveOccurred())
		Expect(archive).To(BeNil())
	})

	It("adds a new archive correctly - ocp versions greater than/ equal to MinimalVersionForNmstatectl", func() {

		clusterProxyInfo := ClusterProxyInfo{
			HTTPProxy:  "http://10.10.1.1:3128",
			HTTPSProxy: "https://10.10.1.1:3128",
			NoProxy:    "quay.io",
		}

		staticnetworkConfigOutput := []staticnetworkconfig.StaticNetworkConfigData{
			{
				FilePath:     "1.yml",
				FileContents: "1.yml contents",
			},
			{
				FilePath:     "2.yml",
				FileContents: "2.yml contents",
			},
		}

		archive, err := RamdiskImageArchive(staticnetworkConfigOutput, &clusterProxyInfo, constants.PreNetworkConfigScriptWithNmstatectl, constants.MinimalISONetworkConfigServiceNmstatectl)
		Expect(err).ToNot(HaveOccurred())

		By("checking that the files are present in the archive")
		gzipReader, err := gzip.NewReader(bytes.NewReader(archive))
		Expect(err).ToNot(HaveOccurred())

		var serviceLink, rootfsServiceConfigContent string
		r := cpio.NewReader(gzipReader)
		for {
			hdr, err := r.Next()
			if err == io.EOF {
				break
			}
			Expect(err).ToNot(HaveOccurred())
			switch hdr.Name {
			case "/etc/1.yml":
				configBytes, err := io.ReadAll(r)
				Expect(err).ToNot(HaveOccurred())
				Expect(string(configBytes)).To(Equal("1.yml contents"))
			case "/etc/2.yml":
				configBytes, err := io.ReadAll(r)
				Expect(err).ToNot(HaveOccurred())
				Expect(string(configBytes)).To(Equal("2.yml contents"))
			case "/usr/local/bin/pre-network-manager-config.sh":
				scriptBytes, err := io.ReadAll(r)
				Expect(string(scriptBytes)).To(Equal(constants.PreNetworkConfigScriptWithNmstatectl))
				Expect(err).ToNot(HaveOccurred())
			case "/etc/systemd/system/pre-network-manager-config.service":
				serviceBytes, err := io.ReadAll(r)
				Expect(string(serviceBytes)).To(Equal(constants.MinimalISONetworkConfigServiceNmstatectl))
				Expect(err).ToNot(HaveOccurred())
			case "/etc/systemd/system/initrd.target.wants/pre-network-manager-config.service":
				Expect(hdr.Mode & cpio.ModeSymlink).ToNot(BeZero())
				serviceLink = hdr.Linkname
			case "/etc/systemd/system/coreos-livepxe-rootfs.service.d/10-proxy.conf":
				rootfsServiceConfigBytes, err := io.ReadAll(r)
				Expect(err).ToNot(HaveOccurred())
				rootfsServiceConfigContent = string(rootfsServiceConfigBytes)
			}
		}

		Expect(serviceLink).To(Equal("/etc/systemd/system/pre-network-manager-config.service"))

		rootfsServiceConfig := fmt.Sprintf("[Service]\n"+
			"Environment=http_proxy=%s\nEnvironment=https_proxy=%s\nEnvironment=no_proxy=%s\n"+
			"Environment=HTTP_PROXY=%s\nEnvironment=HTTPS_PROXY=%s\nEnvironment=NO_PROXY=%s",
			clusterProxyInfo.HTTPProxy, clusterProxyInfo.HTTPSProxy, clusterProxyInfo.NoProxy,
			clusterProxyInfo.HTTPProxy, clusterProxyInfo.HTTPSProxy, clusterProxyInfo.NoProxy)
		Expect(rootfsServiceConfigContent).To(Equal(rootfsServiceConfig))
	})
	It("returns nothing when given nothing - ocp versions greater than/ equal to MinimalVersionForNmstatectl", func() {
		archive, err := RamdiskImageArchive(
			[]staticnetworkconfig.StaticNetworkConfigData{},
			&ClusterProxyInfo{},
			constants.PreNetworkConfigScriptWithNmstatectl, constants.MinimalISONetworkConfigServiceNmstatectl)
		Expect(err).ToNot(HaveOccurred())
		Expect(archive).To(BeNil())
	})

	It("handles authenticated proxy with and without urlencoded special characters", func() {
		testCases := []struct {
			description      string
			clusterProxyInfo ClusterProxyInfo
		}{
			{
				description: "with urlencoded special characters",
				clusterProxyInfo: ClusterProxyInfo{
					HTTPProxy:  "http://user%40example.com:pas%25word%5D@10.10.1.1:3128",
					HTTPSProxy: "https://user%40example.com:pas%25word%5D@10.10.1.1:3128",
					NoProxy:    "quay.io",
				},
			},
			{
				description: "without urlencoded special characters",
				clusterProxyInfo: ClusterProxyInfo{
					HTTPProxy:  "http://userexample.com:pasword@10.10.1.1:3128",
					HTTPSProxy: "https://userexample.com:pasword@10.10.1.1:3128",
					NoProxy:    "quay.io",
				},
			},
		}

		for _, testCase := range testCases {
			tc := testCase

			By("checking proxy settings " + tc.description + "are correctly included in the archive")

			archive, err := RamdiskImageArchive([]staticnetworkconfig.StaticNetworkConfigData{}, &tc.clusterProxyInfo, constants.PreNetworkConfigScriptWithNmstatectl, constants.MinimalISONetworkConfigServiceNmstatectl)
			Expect(err).ToNot(HaveOccurred())

			gzipReader, err := gzip.NewReader(bytes.NewReader(archive))
			Expect(err).ToNot(HaveOccurred())

			var rootfsServiceConfigContent string
			r := cpio.NewReader(gzipReader)
			for {
				hdr, err := r.Next()
				if err == io.EOF {
					break
				}
				Expect(err).ToNot(HaveOccurred())
				if hdr.Name == "/etc/systemd/system/coreos-livepxe-rootfs.service.d/10-proxy.conf" {
					rootfsServiceConfigBytes, err := io.ReadAll(r)
					Expect(err).ToNot(HaveOccurred())
					rootfsServiceConfigContent = string(rootfsServiceConfigBytes)
				}
			}

			expectedRootfsServiceConfig := fmt.Sprintf(
				"[Service]\n"+
					"Environment=http_proxy=%s\nEnvironment=https_proxy=%s\nEnvironment=no_proxy=%s\n"+
					"Environment=HTTP_PROXY=%s\nEnvironment=HTTPS_PROXY=%s\nEnvironment=NO_PROXY=%s",
				strings.ReplaceAll(tc.clusterProxyInfo.HTTPProxy, "%", "%%"), strings.ReplaceAll(tc.clusterProxyInfo.HTTPSProxy, "%", "%%"), tc.clusterProxyInfo.NoProxy,
				strings.ReplaceAll(tc.clusterProxyInfo.HTTPProxy, "%", "%%"), strings.ReplaceAll(tc.clusterProxyInfo.HTTPSProxy, "%", "%%"), tc.clusterProxyInfo.NoProxy)

			Expect(rootfsServiceConfigContent).To(Equal(expectedRootfsServiceConfig))
		}
	})
})
