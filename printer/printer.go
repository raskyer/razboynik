package printer

import (
	"fmt"

	"github.com/fatih/color"
)

const SPACE = "     "

func intro() {
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
	fmt.Print("\n\n")
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

func det_intro(detail string, s string) {
	fmt.Print("\n")
	color.Cyan(SPACE + "### |" + detail + "| ###")
	color.Cyan(SPACE + "-----" + s + "-----")
}

func Start() {
	intro()
	color.Green(SPACE + "### STARTING ###")
	color.Green(SPACE + "----------------")
	color.White(SPACE + "Trying to communicate with server...")
	fmt.Print("\n")
}

func Generating() {
	intro()
	color.Green(SPACE + "### GENERATING ###")
	color.Green(SPACE + "------------------")
	fmt.Print("\n")
}

func SetupError(i int) {
	err_intro()
	color.White("An error occured during configuration")

	if i == 0 {
		color.White("Flag -u (url) is required")
	} else if i == 1 {
		color.White("Method is between 0 (default) and 3.")
		color.White("[0 => GET, 1 => POST, 2 => HEADER, 3 => COOKIE]")
	}
}

func Error(err error) {
	err_intro()
	color.White(SPACE + err.Error())
}

func End() {
	det_intro("BASH", "----")
	color.White(SPACE + "Meterpreter ready !")
	fmt.Print("\n\n\n")
}
