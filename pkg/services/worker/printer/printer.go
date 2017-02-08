package printer

import (
	"fmt"
	"strings"

	"os"

	"github.com/eatbytes/razboynik/pkg/services/worker/gflag"
	"github.com/fatih/color"
)

func PrintIntro() {
	if gflag.Silent {
		return
	}

	fmt.Fprint(os.Stdout, "\n")
	fmt.Fprintln(os.Stdout, color.WhiteString("███████████████████████████████████████████████████████████████████████"))
	fmt.Fprintln(os.Stdout, color.BlueString("███████████████████████████████████████████████████████████████████████"))
	fmt.Fprint(os.Stdout, "\n")
	fmt.Fprintln(os.Stdout, color.RedString(`██████╗  █████╗ ███████╗██████╗  ██████╗ ██╗   ██╗███╗   ██╗██╗██╗  ██╗
██╔══██╗██╔══██╗╚══███╔╝██╔══██╗██╔═══██╗╚██╗ ██╔╝████╗  ██║██║██║ ██╔╝
██████╔╝███████║  ███╔╝ ██████╔╝██║   ██║ ╚████╔╝ ██╔██╗ ██║██║█████╔╝ 
██╔══██╗██╔══██║ ███╔╝  ██╔══██╗██║   ██║  ╚██╔╝  ██║╚██╗██║██║██╔═██╗ 
██║  ██║██║  ██║███████╗██████╔╝╚██████╔╝   ██║   ██║ ╚████║██║██║  ██╗
╚═╝  ╚═╝╚═╝  ╚═╝╚══════╝╚═════╝  ╚═════╝    ╚═╝   ╚═╝  ╚═══╝╚═╝╚═╝  ╚═╝`))
	fmt.Fprintln(os.Stdout, color.BlueString("_______________________________________________________________________"))
	fmt.Fprint(os.Stdout, "\n")

	fmt.Fprintln(os.Stdout, "из России с любовью <3 !")
	fmt.Fprintln(os.Stdout, color.YellowString("version: 2.0.0"))
}

func PrintTitle(str string) {
	var i, lenght int

	if gflag.Silent {
		return
	}

	lenght = len(str)

	fmt.Fprintln(os.Stdout, color.BlueString("\n███ "+strings.ToUpper(str)+" ███"))

	for i = 0; i < lenght+8; i++ {
		fmt.Fprint(os.Stdout, color.BlueString("-"))
	}

	fmt.Fprint(os.Stdout, "\n")
}

func PrintSection(section string, str string) {
	if gflag.Silent {
		return
	}

	PrintTitle(section)
	fmt.Fprintln(os.Stdout, color.WhiteString(str), "\n")
}

func PrintSectionI(section string, i ...interface{}) {
	if gflag.Silent {
		return
	}

	PrintTitle(section)

	for _, item := range i {
		fmt.Fprintln(os.Stdout, item)
	}
}

func PrintlnI(i ...interface{}) {
	fmt.Fprintln(os.Stdout, i)
}

func Println(str string) {
	Print(str + "\n")
}

func Print(str string) {
	fmt.Fprintf(os.Stdout, str)
}

func PrintError(err error) {
	fmt.Fprintln(os.Stderr, color.RedString(err.Error()))
}

func PrintWarning(str string) {
	fmt.Fprintln(os.Stdout, color.YellowString(str))
}
