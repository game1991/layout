package api

import (
	"github.com/spf13/cobra" // for cobra.Command
)

// StartCmd cmd args
var StartCmd = &cobra.Command{

	Use:          "serve",
	Short:        "Start the server",
	Example:      "helloworld serve -d ../configs",
	SilenceUsage: true,
	RunE: func(_ *cobra.Command, _ []string) error {
		return nil
	},
}
