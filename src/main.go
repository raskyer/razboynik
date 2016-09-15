package main

import (
	"io"
	"log"

	"github.com/chzyer/readline"
)

func main() {
	main := CreateMainApp()
	main.Start()

	Global.Main = main

	mainLoop()
}

func mainLoop() {
	main := Global.Main
	prompt := main.GetPrompt()
	defer prompt.Close()
	log.SetOutput(prompt.Stderr())

	for main.IsRunning() {
		line, err := prompt.Readline()
		if err == readline.ErrInterrupt || err == io.EOF {
			return
		}

		if len(line) == 0 {
			continue
		}

		command := main.GetCommand(line)
		main.Run(command)
	}
}

func createBash() {
	bash := CreateBashApp()
	bash.Start()

	Global.Bash = bash

	if bash.IsRunning() {
		bashLoop()
	}
}

func bashLoop() {
	bash := Global.Bash
	prompt := bash.GetPrompt()
	defer prompt.Close()
	log.SetOutput(prompt.Stderr())

	for bash.IsRunning() {
		line, err := prompt.Readline()
		if err == readline.ErrInterrupt || err == io.EOF {
			Global.Stop()
			return
		}

		if len(line) == 0 {
			continue
		}

		bash.Run(line)
	}
}
