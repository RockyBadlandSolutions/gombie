package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os/exec"
	"path"
	"runtime"
	"strings"
)

func getVersionJSON(ver string) VersionMeta {
	verFile, err := ioutil.ReadFile(path.Join(getMinecraftDirectory(), "versions", ver, ver + ".json"))
	if err != nil {
		panic("Unable to get version file")
	}
	var verJSON VersionMeta
	err = json.Unmarshal(verFile, &verJSON)
	if err != nil {
		panic("Unable to decode version JSON: "+err.Error())
	}
	return verJSON
}

func getLibs(ver string) string {
	var classpathSeparator string
	switch runtime.GOOS {
	case "windows":
		classpathSeparator = ";"
	default:
		classpathSeparator = ":"
	}

	var sb = ""
	verMeta := getVersionJSON(ver)
	for _, lib := range verMeta.Libraries {
		if !allowedCheck(lib.Rules) { continue }
		classPath := getClassiferPath(lib)
		sb += path.Join(getMinecraftDirectory(), "libraries", lib.Artifact.Path) + classpathSeparator
		if classPath != "" { sb += path.Join(getMinecraftDirectory(), "libraries", classPath) + classpathSeparator }
	}
	sb += path.Join(getMinecraftDirectory(), "versions", verMeta.ID, verMeta.ID + ".jar")
	if runtime.GOOS == "windows" { return strings.Replace(sb, "/", "\\", -1) }
	return sb
}

func getClassiferPath(library Library) string {
	classifers := library.Classifies
	var classifer string
	switch runtime.GOOS {
	case "windows":
		classifer = classifers.Windows.Path
	case "linux":
		classifer = classifers.Linux.Path
	case "darwin":
		classifer = classifers.Osx.Path
	}
	return classifer
}

func allowedCheck(rules []Rule) bool {
	if len(rules) > 0 {
		var os string
		if runtime.GOOS == "darwin" { os="osx" } else { os=runtime.GOOS }
		for _, rule := range rules {
			if rule.Os.Name == "" { return true }
			if rule.Action == "disallow" && rule.Os.Name == os {
				return false
			} else if rule.Action == "allow" && rule.Os.Name != os{
				return false
			} else {
				return true
			}
		}
	}
	return true
}

func getArgs(ver string) string {
	var verMeta = getVersionJSON(ver)
	args := verMeta.Arguments.Game
	JVMArgs := verMeta.Arguments.Jvm
	sb := ""

	for _, jvmArg := range JVMArgs {
		if !allowedCheck(jvmArg.Rules) || strings.Contains(jvmArg.Values[0], "Windows 10") { continue }
		for _, val := range jvmArg.Values {
			if strings.Contains(val, " ") { val = fmt.Sprintf("\"%v\"", val) }
			sb += setJVMArgs(val, ver) + " "
		}
	}
	sb += verMeta.MainClass + " "
	for _, arg := range args {
		for _, value := range arg.Values {
			sb += setArgs(value, ver) + " "
		}
	}
	return sb
}

func setArgs(arg string, ver string) string  {
	var result string
	switch arg {
	case "${auth_player_name}":
		// FIX USERNAME SETTING
		result = "Player"
	case "${version_name}":
		result = ver
	case "${game_directory}":
		result = getMinecraftDirectory()
	case "${assets_root}":
		result = path.Join(getMinecraftDirectory(), "assets")
	case "${assets_index_name}":
		result = ver
	case "${auth_uuid}":
		// UUID For offline mode
		result = "00000000-0000-0000-0000-000000000000"
	case "${auth_access_token}":
		// Access token for offline mode
		result = "null"
	case "${user_type}":
		result = "legacy"
	case "${version_type}":
		result = getVersionJSON(ver).Type
	case "${resolution_width}":
		result = "640"
	case "${resolution_height}":
		result = "480"
	case "--demo":
		result = ""
	default:
		result = arg
	}
	if runtime.GOOS == "windows" {
		result = strings.Replace(result, "/", "\\", -1)
	}
	return result

}

func setJVMArgs(arg string, ver string) (result string) {
	result = strings.Replace(arg, "${natives_directory}", path.Join(getMinecraftDirectory(), "versions", ver, "natives"), 1)
	result = strings.Replace(result, "${launcher_name}", "GOL", 1)
	result = strings.Replace(result, "${launcher_version}", "v0", 1)
	if runtime.GOOS == "windows" { result=strings.Replace(result, "/", "\\", -1) }
	if arg == "${classpath}" { result = getLibs(ver) }
	return result
}

func Launch(ver string) {
	cmd := exec.Cmd{Path: getJAVAExecutable(), Args: strings.Split(getArgs(ver), " ")}
	err := cmd.Run()
	if err != nil {
		panic("Error while starting game: " + err.Error())
	}
}