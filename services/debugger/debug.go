package debugger

import "github.com/eatbytes/razboynik/services/printer"
import "github.com/eatbytes/razboy"

func HTTP(rzRes *razboy.RESPONSE) {
	printer.PrintSection("Debug", "Debugging HTTP")
	printer.PrintlnI("Request \n", *rzRes.GetRequest().GetHTTP())
	printer.PrintlnI("Response \n", *rzRes.GetHTTP())
}
