package printer

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

func PrintIntro() {
	fmt.Print("\n")
	color.Cyan("########")
	color.Cyan("#       #")
	color.Cyan("#      #")
	color.Cyan("#######         #    #####    #####     ####    #   #   #   #   #   #   #")
	color.Cyan("#     #        # #       #    #    #   #    #    # #    ##  #   #   #  # ")
	color.Cyan("#      #      #####     #     #####    #    #     #     # # #   #   ###  ")
	color.Cyan("#       #     #   #    #      #    #   #    #     #     #  ##   #   #  # ")
	color.Cyan("#        #    #   #   #####   #####     ####      #     #   #   #   #   #")
	color.Magenta("#########################################################################")
	fmt.Print("\n")

	color.White("из России с любовью <3 !")
	fmt.Print("\n")
}

func err_intro() {
	fmt.Print("\n")
	color.Red("### ERROR ###")
	color.Red("-------------")
}

func suc_intro() {
	fmt.Print("\n")
	color.Green("### SUCCESS ###")
	color.Green("---------------")
}

func det_intro(str string) {
	var i, lenght int

	lenght = len(str)

	fmt.Print("\n")
	color.Cyan("### " + strings.ToUpper(str) + " ###")

	for i < lenght+8 {
		f := color.CyanString("-")
		fmt.Print(f)
		i++
	}

	fmt.Print("\n")
}

func PrintStart() {
	color.Green("### STARTING ###")
	color.Green("----------------")
	color.White("Trying to communicate with server...")
	fmt.Print("\n")
}

func PrintError(err error) {
	err_intro()
	color.White(err.Error())
	fmt.Print("\n")
}

func PrintSection(section string, str string) {
	det_intro(section)
	color.White(str)
	fmt.Print("\n")
}

func Println(str string) {
	Print(str + "\n")
}

func Print(str string) {
	fmt.Print(str)
}
