package common

import "github.com/openshift/assisted-service/models"

const (
	VmwareManufacturer string = "VMware, Inc."
)

var supportedManufacturers = map[string]models.PlatformType{
	VmwareManufacturer: models.PlatformTypeVsphere,
}

func AppendPlatformFromVendor(vendor models.SystemVendor, hostsPlatform *[]models.PlatformType) {
	platform := supportedManufacturers[vendor.Manufacturer]
	if platform != "" {
		*hostsPlatform = append(*hostsPlatform, platform)
	}
}

func DecideSupportedPlatforms(hostsPlatforms []models.PlatformType, hostsCount int) *[]models.PlatformType {
	supportedPlatforms := []models.PlatformType{models.PlatformTypeBaremetal}

	// SNO or no hosts
	if hostsCount == 0 || hostsCount == 1 || len(hostsPlatforms) < hostsCount {
		return &supportedPlatforms
	}

	additionalPlatform := hostsPlatforms[0]
	for _, platform := range hostsPlatforms {
		if additionalPlatform != platform {
			return &supportedPlatforms
		}
	}

	supportedPlatforms = append(supportedPlatforms, additionalPlatform)
	return &supportedPlatforms
}
