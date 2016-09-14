package php

import (
	"fmt"
	"net/http"
	"network"
	"strings"
	"worker"

	"github.com/urfave/cli"
)

func Raw(c *cli.Context) {
	cmd := strings.Join(c.Args(), " ")

	network.NET.Send(cmd, phpEnd)
}

func phpEnd(r *http.Response) {
	buffer := worker.GetBody(r)
	base64 := string(buffer)
	fmt.Println(base64)
}
