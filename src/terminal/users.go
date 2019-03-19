package terminal

import (
	"fmt"
	"github.com/spf13/cobra"
	"snoopy/shell"
)

var usersCmd = &shell.Command{
	Command: &cobra.Command{
		Use:   "users",
		Short: "list all users on the remote device",
		Long:  "list all users on the remote device",
	},
	Run: users,
}

func init() {
	remoteShell.AddCommand(usersCmd)
}

func users(sh *shell.Shell, args []string) error {
	fmt.Println("listing all users on this system... root, just root")
	return nil
}
