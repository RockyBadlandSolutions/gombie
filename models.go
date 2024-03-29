package main

import "time"

type launcherMeta struct {
	Latest metaLatest		`json:"latest"`
	Versions []metaVersion	`json:"versions"`
}

type metaLatest struct {
	Release string 		`json:"release"`
	Snapshot string 	`json:"snapshot"`
}

type metaVersion struct {
	Id string			`json:"id"`
	Type string			`json:"type"`
	Url string			`json:"url"`
	Time string			`json:"time"`
	ReleaseTime string	`json:"release_time"`
}

type metaVersionMin struct {
	Id string	`json:"id"`
	Type string	`json:"type"`
}

type AutoGenerated struct {
	Arguments struct {
		Game []interface{} `json:"game"`
		Jvm  []interface{} `json:"jvm"`
	} `json:"arguments"`
	AssetIndex struct {
		ID        string `json:"id"`
		Sha1      string `json:"sha1"`
		Size      int    `json:"size"`
		TotalSize int    `json:"totalSize"`
		URL       string `json:"url"`
	} `json:"assetIndex"`
	Assets    string `json:"assets"`
	Downloads struct {
		Client struct {
			Sha1 string `json:"sha1"`
			Size int    `json:"size"`
			URL  string `json:"url"`
		} `json:"client"`
		ClientMappings struct {
			Sha1 string `json:"sha1"`
			Size int    `json:"size"`
			URL  string `json:"url"`
		} `json:"client_mappings"`
		Server struct {
			Sha1 string `json:"sha1"`
			Size int    `json:"size"`
			URL  string `json:"url"`
		} `json:"server"`
		ServerMappings struct {
			Sha1 string `json:"sha1"`
			Size int    `json:"size"`
			URL  string `json:"url"`
		} `json:"server_mappings"`
	} `json:"downloads"`
	ID        string `json:"id"`
	Libraries []struct {
		Downloads struct {
			Artifact struct {
				Path string `json:"path"`
				Sha1 string `json:"sha1"`
				Size int    `json:"size"`
				URL  string `json:"url"`
			} `json:"artifact"`
		} `json:"downloads"`
		Name  string `json:"name"`
		Rules []struct {
			Action string `json:"action"`
			Os     struct {
				Name string `json:"name"`
			} `json:"os"`
		} `json:"rules,omitempty"`
		Natives struct {
			Osx string `json:"osx"`
		} `json:"natives,omitempty"`
		Extract struct {
			Exclude []string `json:"exclude"`
		} `json:"extract,omitempty"`
	} `json:"libraries"`
	Logging struct {
		Client struct {
			Argument string `json:"argument"`
			File     struct {
				ID   string `json:"id"`
				Sha1 string `json:"sha1"`
				Size int    `json:"size"`
				URL  string `json:"url"`
			} `json:"file"`
			Type string `json:"type"`
		} `json:"client"`
	} `json:"logging"`
	MainClass              string    `json:"mainClass"`
	MinimumLauncherVersion int       `json:"minimumLauncherVersion"`
	ReleaseTime            time.Time `json:"releaseTime"`
	Time                   time.Time `json:"time"`
	Type                   string    `json:"type"`
}

type VersionMeta struct {
	Arguments struct {
		Game []struct {
			Rules []struct {
				Action   string `json:"action"`
				Features struct {
					HasCustomResolution bool `json:"has_custom_resolution"`
					IsDemoUser          bool `json:"is_demo_user"`
				} `json:"features"`
			} `json:"rules"`
			Values []string `json:"values"`
		} `json:"game"`
		Jvm []struct {
			Rules []Rule `json:"rules"`
			Values []string `json:"values"`
		} `json:"jvm"`
	} `json:"arguments"`
	AssetIndex struct {
		ID        string `json:"id"`
		Size      int64  `json:"size"`
		TotalSize int64  `json:"totalSize"`
		URL       string `json:"url"`
	} `json:"assetIndex"`
	Assets    string `json:"assets"`
	Downloads struct {
		Client struct {
			Sha1 string `json:"sha1"`
			Size int64  `json:"size"`
			URL  string `json:"url"`
		} `json:"client"`
		ClientMappings struct {
			Sha1 string `json:"sha1"`
			Size int64  `json:"size"`
			URL  string `json:"url"`
		} `json:"client_mappings"`
		Server struct {
			Sha1 string `json:"sha1"`
			Size int64  `json:"size"`
			URL  string `json:"url"`
		} `json:"server"`
		ServerMappings struct {
			Sha1 string `json:"sha1"`
			Size int64  `json:"size"`
			URL  string `json:"url"`
		} `json:"server_mappings"`
	} `json:"downloads"`
	ID        string `json:"id"`
	Libraries []Library `json:"libraries"`
	MainClass              string `json:"mainClass"`
	MinimumLauncherVersion int64  `json:"minimumLauncherVersion"`
	ReleaseTime            string `json:"releaseTime"`
	Source                 string `json:"source"`
	Time                   string `json:"time"`
	Type                   string `json:"type"`
}

type Library struct {
	Artifact struct {
		Path string `json:"path"`
		Sha1 string `json:"sha1"`
		Size int64  `json:"size"`
		URL  string `json:"url"`
	} 	`json:"artifact"`
	Classifies struct {
		Javadoc struct {
			Path string `json:"path"`
			Sha1 string `json:"sha1"`
			Size int64  `json:"size"`
			URL  string `json:"url"`
		} `json:"javadoc"`
		Linux struct {
			Path string `json:"path"`
			Sha1 string `json:"sha1"`
			Size int64  `json:"size"`
			URL  string `json:"url"`
		} `json:"linux"`
		Osx struct {
			Path string `json:"path"`
			Sha1 string `json:"sha1"`
			Size int64  `json:"size"`
			URL  string `json:"url"`
		} `json:"osx"`
		Sources struct {
			Path string `json:"path"`
			Sha1 string `json:"sha1"`
			Size int64  `json:"size"`
			URL  string `json:"url"`
		} `json:"sources"`
		Windows struct {
			Path string `json:"path"`
			Sha1 string `json:"sha1"`
			Size int64  `json:"size"`
			URL  string `json:"url"`
		} `json:"windows"`
	} `json:"classifies"`
	Extract struct {
		Exclude []string `json:"exclude"`
	} 	`json:"extract"`
	Name    string 			`json:"name"`
	Natives struct {
		Linux   string `json:"linux"`
		Osx     string `json:"osx"`
		Windows string `json:"windows"`
	} 	`json:"natives"`
	Rules []Rule 	`json:"rules"`
}

type Rule struct {
	Action string `json:"action"`
	Os     struct {
		Name string `json:"name"`
	} `json:"os"`
}