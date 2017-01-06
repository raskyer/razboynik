package printer

import (
	"fmt"
	"strings"

	"github.com/eatbytes/razboynik/services/gflags"
	"github.com/fatih/color"
)

func PrintIntro() {
	if gflags.Silent {
		return
	}

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

func PrintTitle(str string) {
	var i, lenght int

	if gflags.Silent {
		return
	}

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

func PrintSection(section string, str string) {
	if gflags.Silent {
		return
	}

	PrintTitle(section)
	color.White(str)
	fmt.Print("\n")
}

func PrintSectionI(section string, i ...interface{}) {
	if gflags.Silent {
		return
	}

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
