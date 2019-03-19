package shell

import "github.com/spf13/cobra"

var exitCmd = &Command{
	Command: &cobra.Command{
		Use: "exit",
	},
	Run: func(sh *Shell, args []string) error {
		sh.TriggerExit()
		return nil
	},
}
