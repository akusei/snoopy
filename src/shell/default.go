package shell

import (
	"fmt"
)

func commandNotFoundHandler(sh *Shell, cmdLine []string) error {
	return fmt.Errorf("%s: command not found", cmdLine[0])
}
