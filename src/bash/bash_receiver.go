package bash

import (
	"fmt"
	"fuzzer"
	"io"
	"net/http"
	"os"
	"strings"
)

func (b *BashInterface) ReceiveCd(result string) {
	body := fuzzer.Decode(result)
	line := strings.TrimSpace(body)

	if line != "" {
		fuzzer.CMD.SetContext(line)
		b.SetPrompt("\033[31m»\033[0m [Bash]:" + line + "$ ")
		fmt.Println(body)
	}
}

func (b *BashInterface) ReceiveUpload(result string) {
	if result == "1" {
		fmt.Println("File succeedly upload")
		return
	}

	fmt.Println("An error occured")
}

func (b *BashInterface) ReceiveDownload(resp *http.Response) {
	out, err := os.Create("output.txt")
	defer out.Close()

	if err != nil {
		panic(err)
	}

	_, err = io.Copy(out, resp.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println("Downloaded successfully")
}
