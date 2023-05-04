package command

import (
	"fmt"

	"github.com/raphaelmb/go-audio-metadata-cli/internal/interfaces"
)

type Parser struct {
	commands []interfaces.Command
}

func NewParser(command []interfaces.Command) *Parser {
	return &Parser{commands: command}
}

func (p *Parser) Parse(args []string) error {
	if len(args) < 1 {
		help()
		return nil
	}
	subcommand := args[0]
	for _, cmd := range p.commands {
		if cmd.Name() == subcommand {
			cmd.ParseFlags(args[1:])
			return cmd.Run()
		}
	}
	return fmt.Errorf("unknown subcommand: %s", subcommand)
}
