package main

package main

import (
	"git.xq5.com/golang/helloworld/cmd/api"

	"github.com/spf13/cobra"
)


func main() {
	rootCMD := &cobra.Command{
		Use:   "awesome",
		Short: "awesome: an awesome toolkit for go instant applications",
		Long: `awesome: an awesome toolkit for go instant applications.`,
	}

	rootCMD.AddCommand(api.StartCmd)
	// rootCMD.AddCommand(genorm.GenORMCmd)
	// rootCMD.AddCommand(cron.StartCronCmd)
	// rootCMD.AddCommand(xxx.StartCmd)

	if err := rootCMD.Execute(); err != nil {
		panic(err)
	}
}
