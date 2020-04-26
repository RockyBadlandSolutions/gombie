package main

import (
	"github.com/pkg/errors"
	"path"
)

func CheckVersion(versionId string) error {
	// Check if version installable
	for _, version := range getInstalledVersions() {
		if version.Id == versionId {
			return errors.Errorf("This version already installed")
		}
	}

	for _, version := range getVersions().Versions {
		if version.Id == versionId {
			checkDir(path.Join(getMinecraftDirectory(), "versions"))
			InstallVersion(versionId)

		}
	}
	return nil

}

func InstallVersion(versionId string) {
	
}