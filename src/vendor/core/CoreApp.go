package core

import (
	"strings"
	"time"

	"github.com/chzyer/readline"
	"github.com/urfave/cli"
)

var Running = true
var RunningMain = true
var RunningBash = false

type CoreInterface struct {
	cmd         *cli.App
	completer   *readline.PrefixCompleter
	commands    []cli.Command
	isConnected bool
}

func Create() *CoreInterface {
	c := CoreInterface{}

	c.commands = []cli.Command{
		generateDefinition,
		cmdDefinition,
		srvDefinition,
		helpDefinition,
		exitDefinition,
	}
	c.completer = readline.NewPrefixCompleter()
	buildCompleter(&c.commands, &c.completer)

	c.cmd = createCli(&c.commands)
	addInformation(c.cmd)

	return &c
}

func buildCompleter(c *[]cli.Command, co **readline.PrefixCompleter) {
	command := *c
	completer := *co
	lgt := len(command)

	for i := 0; i < lgt; i++ {
		pc := readline.PcItem(command[i].Name)
		buildChild(&command[i].Subcommands, &pc)
		completer.SetChildren(append(completer.GetChildren(), pc))
	}
}

func buildChild(c *cli.Commands, pc **readline.PrefixCompleter) {
	command := *c
	completer := *pc
	childLgt := len(*c)

	for x := 0; x < childLgt; x++ {
		child := readline.PcItem(command[x].Name)
		completer.SetChildren(append(completer.GetChildren(), child))

		nChild := command[x].Subcommands
		if len(nChild) > 0 {
			buildChild(&nChild, &child)
		}
	}
}

func addInformation(app *cli.App) {
	app.Name = "Fuzzer"
	app.Version = "1.0.0"
	app.Compiled = time.Now()
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Kamikaze",
			Email: "leakleass@protomail.com",
		},
	}
	app.Copyright = "(c) 2016 Eat Bytes"
	app.Usage = "File upload meterpreter"
}

func createCli(commands *[]cli.Command) *cli.App {
	client := cli.NewApp()
	client.Commands = *commands
	client.CommandNotFound = func(c *cli.Context, command string) {
		cli.ShowAppHelp(c)
	}

	return client
}

func (c *CoreInterface) GetCommand(line string) []string {
	line = strings.TrimSpace(line)
	appName := []string{"Fuzzer"}
	args := strings.Fields(line)
	command := append(appName, args...)

	return command
}

func (c *CoreInterface) GetPrompt() *readline.Instance {
	l, err := readline.NewEx(&readline.Config{
		Prompt:          "\033[31m»\033[0m ",
		HistoryFile:     "/tmp/readline.tmp",
		AutoComplete:    c.completer,
		InterruptPrompt: "^C",
		EOFPrompt:       "exit",
	})

	if err != nil {
		panic(err)
	}

	return l
}

func (c *CoreInterface) Run(command []string) {
	c.cmd.Run(command)
}
