package build

import (
	"strconv"
	"time"
)

// Version is the app-global version string, which should be substituted with a
// real value during build.
var Version = "UNKNOWN"

// Name is the name of the compiled software.
var Name = ""
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

func BuildInfo() *SysInfo {
	parseInt, _ := strconv.ParseInt(LastTime, 10, 64)
	tm := time.Unix(parseInt, 0)
	SystemInfo = &SysInfo{
		Branch:      Branch,
		Tag:         Tag,
		LastTime:    tm.Format("2006-01-02 15:04:05"),
		LastCommit:  LastCommit,
		ServiceName: Name,
		GitRepoPATH: GitRepoPATH,
	}
	return SystemInfo
}
