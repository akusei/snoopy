package shell

type OnExitHandler func(*Shell) error

type CommandHandler func(sh *Shell, args []string) error

type CommandNotFoundHandler func(sh *Shell, cmdLine []string) error
