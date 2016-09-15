package main

import (
	"io"
	"log"

	"github.com/chzyer/readline"
)

func main() {
	main := CreateMainApp()
	main.Start()

	mainLoop(main)
}

func mainLoop(main *MainInterface) {
	prompt := main.GetPrompt()
	defer prompt.Close()
	log.SetOutput(prompt.Stderr())

	for main.IsRunning() {
		line, err := prompt.Readline()
		if err == readline.ErrInterrupt || err == io.EOF {
			break
		}

		if len(line) == 0 {
			continue
		}

		command := main.GetCommand(line)
		main.Run(command)
	}

	if Global.BashSession {
		bash := CreateBashApp()
		bash.Start()

		if bash.IsRunning() {
			bashLoop(bash, main)
		} else {
			Global.MainSession = true
			main.Start()
			mainLoop(main)
		}
	}
}

func bashLoop(bash *BashInterface, parent *MainInterface) {
	prompt := bash.GetPrompt()
	defer prompt.Close()
	log.SetOutput(prompt.Stderr())

	for bash.IsRunning() {
		line, err := prompt.Readline()
		if err == readline.ErrInterrupt || err == io.EOF {
			bash.Stop()
		}

		if len(line) == 0 {
			continue
		}

		bash.Run(line)
	}

	if Global.MainSession {
		mainLoop(parent)
	}
}
