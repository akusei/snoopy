package shell

// TODO: correct issue when doing help --help and then help again
// TODO: add aliases map so I can properly show help

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

type Shell struct {
	commands       map[string]*Command
	aliases        map[string]*Command
	defaultCommand CommandNotFoundHandler
	onExit         OnExitHandler

	target     string
	prompt     string
	username   string
	workingDir string
	exitShell  bool
}

func New(opts ...Option) (*Shell, error) {
	newShell := &Shell{
		target:     "127.0.0.1",
		prompt:     "$",
		username:   "user",
		workingDir: "~",
		commands:   make(map[string]*Command),
		aliases:    make(map[string]*Command),
	}

	_ = WithCommandNotFound(commandNotFoundHandler)(newShell)
	newShell.AddCommand(exitCmd)
	newShell.AddCommand(helpCmd)

	for _, opt := range opts {
		if err := opt(newShell); err != nil {
			return nil, err
		}
	}

	return newShell, nil
}

func (s *Shell) Run(target string) error {
	s.target = target

	for s.exitShell = false; !s.exitShell; {
		fmt.Printf("%s@%s:%s%s ", s.username, s.target, s.workingDir, s.prompt)
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		text = strings.TrimRight(text, "\r\n")
		if len(text) == 0 {
			continue
		}

		args := strings.Fields(text)
		cmd := s.findCommand(args[0])

		if cmd == nil {
			if s.defaultCommand != nil {
				_ = s.defaultCommand(s, args)
			}
			continue
		}

		cmd.ResetFlags()
		cmd.SetArgs(args[1:])
		_ = cmd.Execute()
	}

	var err error
	if s.onExit != nil {
		err = s.onExit(s)
	}

	return err
}

func (s *Shell) TriggerExit() {
	s.exitShell = true
}

func (s *Shell) AddCommand(cmd *Command) {
	s.commands[cmd.Command.Name()] = cmd
	for _, alias := range cmd.Command.Aliases {
		s.aliases[alias] = cmd
	}

	if cmd.Run == nil {
		return
	}

	cmd.Command.RunE = func(c *cobra.Command, args []string) error {
		return cmd.Run(s, cmd, args)
	}
}

func (s *Shell) findCommand(name string) *Command {
	if cmd, ok := s.commands[name]; ok {
		return cmd
	}

	if cmd, ok := s.aliases[name]; ok {
		return cmd
	}

	return nil
}

func (s *Shell) User() string {
	return s.username
}

func (s *Shell) Dir() string {
	return s.workingDir
}

func (s *Shell) Prompt() string {
	return s.prompt
}

func (s *Shell) Target() string {
	return s.target
}

func (s *Shell) SetUser(user string) {
	s.username = user
}

func (s *Shell) SetDir(dir string) {
	s.workingDir = dir
}

func (s *Shell) SetPrompt(prompt string) {
	s.prompt = prompt
}

func (s *Shell) SetTarget(target string) {
	s.target = target
}
