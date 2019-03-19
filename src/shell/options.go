package shell

type Option func(sh *Shell) error

func WithPrompt(prompt string) Option {
	return func(sh *Shell) (err error) {
		sh.prompt = prompt
		return
	}
}

func WithUser(name string) Option {
	return func(sh *Shell) (err error) {
		sh.username = name
		return
	}
}

func WithDir(dir string) Option {
	return func(sh *Shell) (err error) {
		sh.workingDir = dir
		return
	}
}

func WithOnExit(exit OnExitHandler) Option {
	return func(sh *Shell) (err error) {
		sh.onExit = exit
		return
	}
}

func WithCommandNotFound(handler CommandNotFoundHandler) Option {
	return func(sh *Shell) (err error) {
		sh.defaultCommand = handler
		return
	}
}
