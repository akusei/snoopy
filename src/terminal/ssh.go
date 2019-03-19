package terminal

import (
	"fmt"
	"github.com/spf13/cobra"
	"snoopy/shell"
)

var sshCmd = &shell.Command{
	Command: &cobra.Command{
		Use:   "ssh",
		Short: "enable/disable ssh",
		Long:  "enable/disable ssh",
		Args:  cobra.ExactArgs(2),
	},
	Run: ssh,
}

func init() {
	remoteShell.AddCommand(sshCmd)
}

func ssh(sh *shell.Shell, args []string) error {
	fmt.Println("ssh is enabled?")
	return nil
}
