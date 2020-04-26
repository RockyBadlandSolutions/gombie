package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"runtime"
)

// Directory to install minecraft
func getMinecraftDirectory() string {
	switch runtime.GOOS {
	case "windows":
		return path.Join(os.Getenv("APPDATA"), "."+minecraftDir)
	case "darwin":
		usr, err := os.UserHomeDir()
		if err != nil {
			panic("can't find user home directory: " + err.Error())
		}
		return path.Join(usr, "Library", "Application Support", minecraftDir)
	default:
		usr, err := os.UserHomeDir()
		if err != nil {
			panic("can't find user home directory: " + err.Error())
		}
		return path.Join(usr, "."+minecraftDir)
	}
}

// Get all minecraft versions from Mojang site
func getVersions() launcherMeta {
	resp, err := http.Get("https://launchermeta.mojang.com/mc/game/version_manifest.json")
	if err != nil {
		log.Print("Unable to get versions info: " + err.Error())
	}
	defer func() {
		err = resp.Body.Close()
		if err != nil {
			log.Print("Unable to get versions info: " + err.Error())
		}
	}()
	var metaStruct launcherMeta
	bbody, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(bbody, &metaStruct)
	return metaStruct
}

// Get all installed versions
func getInstalledVersions() []metaVersionMin {
	mpath := getMinecraftDirectory()
	installDir, err := ioutil.ReadDir(path.Join(mpath, "versions"))
	if err != nil {
		log.Print("Unable to get versions dir")
	}
	var versions []metaVersionMin
	for _, ver := range installDir {
		versionFile := path.Join(mpath, "versions", ver.Name(), ver.Name()+".json")
		body, err := ioutil.ReadFile(versionFile)
		if err != nil {
			panic("Unable to get installed versions: "+err.Error())
		}
		var verInstance metaVersionMin
		err = json.Unmarshal(body, &verInstance)
		if err != nil {
			panic("Unable to get installed versions: "+err.Error())
		}
		versions = append(versions, verInstance)
	}
	return versions

}

// Get all not installed versions
func getAvailableVersions() []metaVersion {
	var versions []metaVersion
	for _, version := range getVersions().Versions {
		for _, installed := range getInstalledVersions() {
			if installed.Id == version.Id && installed.Type == version.Type { continue } else { versions = append(versions, version) }
		}
	}
	return versions
}

// Get path to JAVA.exe file(if available)
func getJAVAExecutable () string {
	switch runtime.GOOS {
	case "windows":
		javaePath := "C:\\Program Files (x86)\\Common Files\\Oracle\\Java\\javapath\\java.exe"
		if _, err := os.Stat(javaePath); err == nil {
			return javaePath
		} else {
			return "java"
		}
	default:
		return "java"



	}
}

// Create dir if not exist
func checkDir(directory string) {
	_, err := os.Stat(directory)
	if os.IsNotExist(err) {
		err = os.MkdirAll(directory, 0755)
		if err != nil {
			log.Fatal("FATAL: Unable to create directories: "+ err.Error())
		}
	}
}


