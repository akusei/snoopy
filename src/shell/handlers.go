package shell

type OnExitHandler func(*Shell) error

type CommandHandler func(sh *Shell, cmd *Command, args []string) error

type CommandNotFoundHandler func(sh *Shell, cmdLine []string) error
