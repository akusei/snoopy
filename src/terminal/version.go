package terminal

import (
	"fmt"
	"github.com/spf13/cobra"
	"snoopy/shell"
)

var versionCmd = &shell.Command{
	Command: &cobra.Command{
		Use:     "version",
		Short:   "show version of remote device",
		Long:    "Display the remote device version",
		Aliases: []string{"ver"},
	},
	Run: version,
}

func init() {
	remoteShell.AddCommand(versionCmd)
}

func version(sh *shell.Shell, args []string) error {
	fmt.Println("FROM VERSION CMD")
	return nil
}
