package main

import (
	"fmt"
	"fuzzer"
)

func (b *BashInterface) Start() {
	if !fuzzer.NET.IsSetup() {
		fmt.Println("You haven't setup the required information, please refer to srv config")
		b.Stop()

		return
	}

	b.running = true
}

func (b *BashInterface) Stop() {
	b.running = false
	fuzzer.CMD.SetContext("")
}

func (b *BashInterface) IsRunning() bool {
	return b.running
}
