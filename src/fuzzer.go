package main

import (
	"bash"
	"io"
	"log"

	"core"

	"github.com/chzyer/readline"
)

func main() {
	mainLoop()
}

func mainLoop() {
	app := core.Create()
	prompt := app.GetPrompt()

	defer prompt.Close()
	log.SetOutput(prompt.Stderr())

	running := &core.Running
	runningMain := &core.RunningMain
	runningBash := &core.RunningBash

	for *running && *runningMain {
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

	if *runningBash {
		bashLoop(runningBash)
	}
}

func bashLoop(running *bool) {
	bash := bash.BSH
	prompt := bash.GetPrompt()

	defer prompt.Close()
	log.SetOutput(prompt.Stderr())

	for *running {
		line, err := prompt.Readline()
		if err == readline.ErrInterrupt || err == io.EOF {
			*running = false
		}

		if len(line) == 0 {
			continue
		}

		bash.Run(line)
	}

	mainLoop()
}
