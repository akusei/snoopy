package terminal

import (
	"fmt"
	"github.com/spf13/cobra"
	"snoopy/shell"
)

var bypassCmd = &shell.Command{
	Command: &cobra.Command{
		Use:   "bypass",
		Short: "generate an auth bypass URL",
		Long:  "creates an authentication bypass page",
	},
	Run: authBypass,
}

func init() {
	remoteShell.AddCommand(bypassCmd)
}

func authBypass(sh *shell.Shell, cmd *shell.Command, args []string) error {
	fmt.Println("Super cool authentication bypass url goes here!")
	return nil
}
