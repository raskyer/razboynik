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

	color.Yellow("из России с любовью <3 !")
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

func det_intro(str string) {
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
	det_intro(section)
	color.White(str)
	fmt.Print("\n")
}

func PrintSectionI(section string, i ...interface{}) {
	det_intro(section)
	fmt.Println(i)
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