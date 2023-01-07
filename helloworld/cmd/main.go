package main

import (
	"helloworld/cmd/api"
	"helloworld/cmd/genorm"

	"github.com/spf13/cobra"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string
	// Version is the version of the compiled software.
	Version string
	// // flagconf is the config flag.
	// flagconf string

	// id, _ = os.Hostname()
)

func main() {
	rootCMD := &cobra.Command{
		Use:   "[project-name]helloworld",
		Short: "A generator for Cobra based Applications",
		Long: `Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	}

	rootCMD.AddCommand(api.StartCmd)
	rootCMD.AddCommand(genorm.GenORMCmd)
	// rootCMD.AddCommand(cron.StartCronCmd)
	// rootCMD.AddCommand(xxx.StartCmd)

	if err := rootCMD.Execute(); err != nil {
		panic(err)
	}
}
