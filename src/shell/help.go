package shell

import (
	"fmt"
	"github.com/spf13/cobra"
)

var helpCmd = &Command{
	Command: &cobra.Command{
		Use:     "help",
		Short:   "This cruft",
		Aliases: []string{"?", "man"},
	},
	Run: showHelp,
}

var hitCount = -1
var mixedUpMotherGoose = []string{
	"Yep",
	"Yep, yeeeeeeep",
	"You're not crazy, they're not in the same order",
	"It's a bug and it'll get fixed",
}

func showHelp(sh *Shell, args []string) error {
	fmt.Println("Available Commands:")

	for _, cmd := range sh.commands {
		fmt.Printf("  %-20v    %s\n", cmd.Name(), cmd.Short)
	}

	if hitCount >= 0 {
		fmt.Println()
		fmt.Println(mixedUpMotherGoose[hitCount])
		if hitCount < len(mixedUpMotherGoose)-1 {
			hitCount++
		}
	} else {
		hitCount++
	}

	fmt.Println()

	return nil
}
