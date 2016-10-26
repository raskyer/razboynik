package printer

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

const SPACE = "     "

func Intro() {
	fmt.Print("\n")
	color.Cyan(SPACE + "########")
	color.Cyan(SPACE + "#")
	color.Cyan(SPACE + "#")
	color.Cyan(SPACE + "#######  #   #   ####    ####    #####   #####")
	color.Cyan(SPACE + "#        #   #      #       #    #       #    #")
	color.Cyan(SPACE + "#        #   #     #       #     #####   #####")
	color.Cyan(SPACE + "#        #   #    #       #      #       #   #")
	color.Cyan(SPACE + "#        #####   #####   #####   #####   #    #")
	color.Magenta(SPACE + "################################################")
	fmt.Print("\n")

	color.White(SPACE + "Hacking web server thanks to php backdoor!")
	fmt.Print("\n")
}

func err_intro() {
	fmt.Print("\n")
	color.Red(SPACE + "### ERROR ###")
	color.Red(SPACE + "-------------")
}

func suc_intro() {
	fmt.Print("\n")
	color.Green(SPACE + "### SUCCESS ###")
	color.Green(SPACE + "---------------")
}

func det_intro(str string) {
	var i, lenght int

	lenght = len(str)

	fmt.Print("\n")
	color.Cyan(SPACE + "### " + strings.ToUpper(str) + " ###")
	fmt.Print(SPACE)

	for i < lenght+8 {
		f := color.CyanString("-")
		fmt.Print(f)
		i++
	}

	fmt.Print("\n")
}

func Start() {
	color.Green(SPACE + "### STARTING ###")
	color.Green(SPACE + "----------------")
	color.White(SPACE + "Trying to communicate with server...")
	fmt.Print("\n")
}

func Generating() {
	color.Green(SPACE + "### GENERATING ###")
	color.Green(SPACE + "------------------")
	fmt.Print("\n")
}

func Error(err error) {
	err_intro()
	color.White(SPACE + err.Error())
	fmt.Print("\n")
}

func PrintSection(section string, str string) {
	det_intro(section)
	color.White(SPACE + str)
	fmt.Print("\n")
}

func Println(str string) {
	Print(str)
	fmt.Print("\n")
}

func Print(str string) {
	fmt.Print(SPACE + str)
}

func Read() (string, error) {
	var input string

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		return input, err
	}

	return input, nil
}

func ReadInt() (int, error) {
	var s string
	var input int

	s, err := Read()
	s = strings.TrimSpace(s)

	if s != "" {
		input, err = strconv.Atoi(s)
	}

	return input, err
}
