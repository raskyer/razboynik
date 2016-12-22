package kernel

import "strings"

type KernelLine struct {
	name string
	raw  string
	arr  []string
	out  string
	err  string
}

func CreateLine(raw string) *KernelLine {
	var (
		arr            []string
		out, err, name string
	)

	out = "&1"
	err = "&2"

	arr = strings.Fields(raw)

	if len(arr) > 0 {
		name = arr[0]
		arr = append(arr[1:], arr[len(arr):]...)
	}

	if i := extractIn(arr, "->"); i != -1 {
		out = arr[i+1]
		arr = append(arr[:i], arr[i+2:]...)
		raw = strings.Join(arr, " ")
	}

	if i := extractIn(arr, "-2>"); i != -1 {
		err = arr[i+1]
		arr = append(arr[:i], arr[i+2:]...)
		raw = strings.Join(arr, " ")
	}

	return &KernelLine{
		name: name,
		raw:  raw,
		arr:  arr,
		out:  out,
		err:  err,
	}
}

func (kl KernelLine) GetRaw() string {
	return kl.raw
}

func (kl KernelLine) GetArr() []string {
	return kl.arr
}

func (kl KernelLine) GetStdout() string {
	return kl.out
}

func (kl KernelLine) GetStderr() string {
	return kl.err
}

func extractIn(slice []string, value string) int {
	for p, v := range slice {
		if v == value {
			return p
		}
	}

	return -1
}
