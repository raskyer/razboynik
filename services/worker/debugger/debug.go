package debugger

import (
	"fmt"
	"net/http/httputil"

	"github.com/eatbytes/razboy"
	"github.com/eatbytes/razboynik/services/printer"
	"github.com/fatih/color"
)

func HTTP(res *razboy.RESPONSE) {
	printer.PrintSection("Debug", "Debugging HTTP")

	color.Yellow("--- Request ---")
	b, _ := httputil.DumpRequestOut(res.GetRequest().GetHTTP(), false)
	fmt.Println(string(b))
	fmt.Println(string(res.GetRequest().GetBody()))

	color.Yellow("--- Respone ---")
	b, _ = httputil.DumpResponse(res.GetHTTP(), true)
	fmt.Println(string(b))
}
