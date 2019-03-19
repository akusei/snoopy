package terminal

import (
	"fmt"
	"snoopy/shell"
)

func commandNotFound(sh *shell.Shell, cmdLine []string) error {
	fmt.Println("not found: sending command to remote server")
	return nil
}
