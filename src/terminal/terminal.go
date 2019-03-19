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

func Run(target string) error {
	return remoteShell.Run(target)
}
