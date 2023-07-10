package new

import "github.com/spf13/cobra"

var CmdNew = &cobra.Command{
	Use:   "new",
	Short: "Create a service template",
	Long:  "Create a service template like helloworld which contains http and grpc web framework",
	Run:   run,
}

func run(cmd *cobra.Command, args []string) {
	
}
