package services

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func Read() (string, error) {
	var input string

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		return input, err
	}

	return input, nil
}

func ReadInt() (int, error) {
	var s string
	var input int

	s, err := Read()
	s = strings.TrimSpace(s)

	if s != "" {
		input, err = strconv.Atoi(s)
	}

	return input, err
}
