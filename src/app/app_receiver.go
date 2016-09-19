package app

import "fmt"

func (main *MainInterface) ReceiveUpload(result string) {
	if result == "1" {
		fmt.Println("File succeedly upload")
		return
	}

	fmt.Println("An error occured")
}

func (main *MainInterface) ReceiveTest(str string) {
	if str == "1" {
		fmt.Println("Connected")
		return
	}

	fmt.Println("Error")
}
