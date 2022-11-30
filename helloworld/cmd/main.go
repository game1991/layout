package main

import (
	"helloworld/cmd/api"

	"github.com/spf13/cobra"
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
	// rootCMD.AddCommand(genorm.GenORMCmd)
	// rootCMD.AddCommand(cron.StartCronCmd)
	// rootCMD.AddCommand(xxx.StartCmd)

	if err := rootCMD.Execute(); err != nil {
		panic(err)
	}
}
