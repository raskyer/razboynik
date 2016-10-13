package bash

import "strings"

func Parse(str string) []string {
	strArr := strings.Fields(str)

	if len(strArr) < 2 {
		return nil
	}

	strArr = append(strArr[1:], strArr[len(strArr):]...)

	return strArr
}

func ParseStr(str string) (string, error) {
	strArr := Parse(str)

	if strArr == nil {
		err := new(error)
		return "", *err
	}

	return strings.Join(strArr, " "), nil
}
