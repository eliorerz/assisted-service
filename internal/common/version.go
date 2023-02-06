package common

import (
	"strings"

	"github.com/hashicorp/go-version"
)

func VersionGreaterOrEqual(version1, version2 string) (bool, error) {
	v1, err := version.NewVersion(version1)
	if err != nil {
		return false, err
	}
	v2, err := version.NewVersion(version2)
	if err != nil {
		return false, err
	}

	return v1.GreaterThanOrEqual(v2), nil
}

func BaseVersionGreaterOrEqual(version, versionMayGraterThan string) (bool, error) {
	// return version >= versionMayGraterThan
	version = strings.Split(version, "-")[0]
	versionMayGraterThan = strings.Split(versionMayGraterThan, "-")[0]

	return VersionGreaterOrEqual(version, versionMayGraterThan)
}
