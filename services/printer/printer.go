package printer

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

func PrintIntro() {
	fmt.Print("\n")
	color.White("███████████████████████████████████████████████████████████████████████")
	color.Blue("███████████████████████████████████████████████████████████████████████")
	fmt.Print("\n")
	color.Red(`██████╗  █████╗ ███████╗██████╗  ██████╗ ██╗   ██╗███╗   ██╗██╗██╗  ██╗
██╔══██╗██╔══██╗╚══███╔╝██╔══██╗██╔═══██╗╚██╗ ██╔╝████╗  ██║██║██║ ██╔╝
██████╔╝███████║  ███╔╝ ██████╔╝██║   ██║ ╚████╔╝ ██╔██╗ ██║██║█████╔╝ 
██╔══██╗██╔══██║ ███╔╝  ██╔══██╗██║   ██║  ╚██╔╝  ██║╚██╗██║██║██╔═██╗ 
██║  ██║██║  ██║███████╗██████╔╝╚██████╔╝   ██║   ██║ ╚████║██║██║  ██╗
╚═╝  ╚═╝╚═╝  ╚═╝╚══════╝╚═════╝  ╚═════╝    ╚═╝   ╚═╝  ╚═══╝╚═╝╚═╝  ╚═╝`)
	color.Blue("_______________________________________________________________________")
	fmt.Print("\n")

	color.White("из России с любовью <3 !")
	color.Yellow("version: 2.0.0")
	fmt.Print("\n")
}

func err_intro() {
	fmt.Print("\n")
	color.Red("███ ERROR ███")
	color.Red("_____________")
}

func suc_intro() {
	fmt.Print("\n")
	color.Green("███ SUCCESS ███")
	color.Green("_______________")
}

func PrintTitle(str string) {
	var i, lenght int

	lenght = len(str)

	fmt.Print("\n")
	color.Blue("███ " + strings.ToUpper(str) + " ███")

	for i < lenght+8 {
		f := color.BlueString("-")
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
	PrintTitle(section)
	color.White(str)
	fmt.Print("\n")
}

func PrintSectionI(section string, i ...interface{}) {
	PrintTitle(section)

	for _, item := range i {
		fmt.Println(item)
	}
}

func PrintlnI(i ...interface{}) {
	fmt.Println(i)
}

func Println(str string) {
	Print(str + "\n")
}

func Print(str string) {
	fmt.Print(str)
}
