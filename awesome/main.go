package main

import (
	"fmt"

	"github.com/dimiro1/banner"
	"github.com/game1991/layout/awesome/internal/install"
	"github.com/game1991/layout/awesome/internal/new"
	"github.com/mattn/go-colorable"

	"github.com/spf13/cobra"
)

func init() {
	var tpl = fmt.Sprintln()

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
