package services

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

const SPACE = "     "

func PrintIntro() {
	fmt.Print("\n")
	color.Cyan(SPACE + "########")
	color.Cyan(SPACE + "#       #")
	color.Cyan(SPACE + "#      #")
	color.Cyan(SPACE + "#######         #    #####    #####     ####    #   #   #   #   #   #   #")
	color.Cyan(SPACE + "#     #        # #       #    #    #   #    #    # #    ##  #   #   #  # ")
	color.Cyan(SPACE + "#      #      #####     #     #####    #    #     #     # # #   #   ###  ")
	color.Cyan(SPACE + "#       #     #   #    #      #    #   #    #     #     #  ##   #   #  # ")
	color.Cyan(SPACE + "#        #    #   #   #####   #####     ####      #     #   #   #   #   #")
	color.Magenta(SPACE + "#########################################################################")
	fmt.Print("\n")

	color.White(SPACE + "из России с любовью <3 !")
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

func PrintStart() {
	color.Green(SPACE + "### STARTING ###")
	color.Green(SPACE + "----------------")
	color.White(SPACE + "Trying to communicate with server...")
	fmt.Print("\n")
}

func PrintError(err error) {
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
	Print(str + "\n")
}

func Print(str string) {
	fmt.Print(SPACE + str)
}
