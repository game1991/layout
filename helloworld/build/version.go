package build

// Version is the app-global version string, which should be substituted with a
// real value during build.
var Version = "UNKNOWN"
var Branch = ""
var Tag = ""
var LastTime = ""
var LastCommit = ""
var GitRepoPATH = ""

var SystemInfo *SysInfo

type SysInfo struct {
	Branch      string `json:"branch,omitempty"`
	Tag         string `json:"tag,omitempty"`
	LastTime    string `json:"lastTime"`
	LastCommit  string `json:"lastCommit"`
	ServiceName string `json:"serviceName"`
	GitRepoPATH string `json:"gitRepoPath"`
}
