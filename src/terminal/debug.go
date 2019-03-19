package terminal

import (
	"fmt"
	"github.com/spf13/cobra"
	"snoopy/shell"
)

var debugCmd = &shell.Command{
	Command: &cobra.Command{
		Use:   "debug",
		Short: "toggle debug output",
		Long:  "toggle shell debug output",
	},
	Run: toggleDebug,
}

func init() {
	remoteShell.AddCommand(debugCmd)
}

func toggleDebug(sh *shell.Shell, args []string) error {
	fmt.Println("might remove this command?")
	return nil
}
