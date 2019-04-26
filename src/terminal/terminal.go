package terminal

import "snoopy/shell"

// initializing this outside of init to ensure its creation
// before other init funcs are called
var remoteShell, _ = shell.New(
	shell.WithPrompt("#"),
	shell.WithCommandNotFound(commandNotFound),
	shell.WithUser("root"),
	shell.WithOnExit(onExit),
	shell.WithDir("/var/www/cgi-bin"),
)

var targetPort uint16
var targetSchema string

func Run(schema string, target string, port uint16) error {
	targetSchema = schema
	targetPort = port
	return remoteShell.Run(target)
}
