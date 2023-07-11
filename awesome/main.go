package main

import (
	"github.com/dimiro1/banner"
	"github.com/game1991/layout/awesome/internal/install"
	"github.com/game1991/layout/awesome/internal/new"
	"github.com/mattn/go-colorable"

	"github.com/spf13/cobra"
)

func init() {
	var tpl = `{{ .Title "awesome" "larry3d" 0 }}
	{{ .AnsiColor.BrightCyan }}The title will be ascii and indented 2 spaces
	{{ .AnsiColor.BrightYellow }}~Welcome to my awesome project~
	{{ .AnsiColor.Magenta }}
	GoVersion: {{ .GoVersion }}
	GOOS: {{ .GOOS }}
	GOARCH: {{ .GOARCH }}
	NumCPU: {{ .NumCPU }}
	GOPATH: {{ .GOPATH }}
	GOROOT: {{ .GOROOT }}
	Compiler: {{ .Compiler }}
	ENV: {{ .Env "GOPATH" }}
	Now: {{ .Now "Monday, 2 Jan 2006" }}
	{{ .AnsiColor.BrightGreen }}This text will appear in Green
	{{ .AnsiColor.BrightRed }}This text will appear in Red
	{{ .AnsiColor.Default }}
	`

	banner.InitString(colorable.NewColorableStdout(), true, true, tpl)
}

func main() {
	rootCMD := &cobra.Command{
		Use:   "awesome",
		Short: "awesome: an awesome toolkit for go instant applications",
		Long:  `awesome: an awesome toolkit for go instant applications.`,
	}

	rootCMD.AddCommand(new.CmdNew)
	rootCMD.AddCommand(install.CmdInstall)
	// rootCMD.AddCommand(cron.StartCronCmd)
	// rootCMD.AddCommand(xxx.StartCmd)

	if err := rootCMD.Execute(); err != nil {
		panic(err)
	}
}
