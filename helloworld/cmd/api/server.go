package api

import (
	"context"
	"log"

	"github.com/spf13/cobra" // for cobra.Command
)

// StartCmd cmd args
var StartCmd = &cobra.Command{

	Use:          "serve",
	Short:        "Run the gRPC hello-world server",
	Example:      "cmd serve -d ../configs",
	SilenceUsage: true,
	RunE: func(_ *cobra.Command, _ []string) error {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Recover error : %v", err)
			}
		}()

		return nil
	},
}

// APP ...
type APP struct {
	ctx    context.Context
	cancle func()
}
