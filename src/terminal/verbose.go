package terminal

import (
	"fmt"
	"github.com/spf13/cobra"
	"snoopy/shell"
)

var verboseCmd = &shell.Command{
	Command: &cobra.Command{
		Use:   "verbose",
		Short: "toggle verbose output",
		Long:  "toggle verbose output",
	},
	Run: verbose,
}

func init() {
	remoteShell.AddCommand(verboseCmd)
}

func verbose(sh *shell.Shell, cmd *shell.Command, args []string) error {
	fmt.Println("I might remove this")
	return nil
}
