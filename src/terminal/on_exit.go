package terminal

import (
	"fmt"
	"snoopy/shell"
)

func onExit(sh *shell.Shell) error {
	fmt.Println("WOOHOOO OnExit, undo things that need undoing...")
	return nil
}
