package shell

import "github.com/spf13/cobra"

type Command struct {
	*cobra.Command
	Run CommandHandler
}

func (cm *Command) AddCommand(shell *Shell, cmd ...*Command) {
	for _, command := range cmd {
		command.Command.RunE = func(c *cobra.Command, args []string) error {
			return command.Run(shell, args)
		}

		cm.Command.AddCommand(command.Command)
	}
}
