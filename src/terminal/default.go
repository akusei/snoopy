package terminal

import (
	"fmt"
	"snoopy/shell"
)

func commandNotFound(sh *shell.Shell, cmdLine []string) error {
	fmt.Println("not found: sending command to remote server")

	//server := fmt.Sprintf("%s://%s:%d", targetSchema, sh.Target(), targetPort)
	//shell, _ := c2.New(server, "test")
	//
	//cmd := commands.New("test", commands.WithArg("arg1", "value1"))
	//shell.Do(cmd)

	return nil
}
