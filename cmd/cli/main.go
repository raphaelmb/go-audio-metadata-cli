package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/raphaelmb/go-audio-metadata-cli/cmd/cli/command"
	"github.com/raphaelmb/go-audio-metadata-cli/internal/interfaces"
)

func main() {
	client := &http.Client{}
	cmds := []interfaces.Command{
		// command.NewGetCommand(client),
		command.NewUploadCommand(client),
		// command.NewListCommand(client),
	}

	parser := command.NewParser(cmds)
	if err := parser.Parse(os.Args[1:]); err != nil {
		os.Stderr.WriteString(fmt.Sprintf("error: %v", err.Error()))
		os.Exit(1)
	}
}
