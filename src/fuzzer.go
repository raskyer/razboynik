package main

import (
	"io"
	"log"

	"core"

	"github.com/chzyer/readline"
)

func main() {
	app := core.Create()
	prompt := app.GetPrompt()

	defer prompt.Close()
	log.SetOutput(prompt.Stderr())

	running := &core.Running

	for *running {
		line, err := prompt.Readline()
		if err == readline.ErrInterrupt || err == io.EOF {
			break
		}

		if len(line) == 0 {
			continue
		}

		command := app.GetCommand(line)
		app.Run(command)
	}
}
