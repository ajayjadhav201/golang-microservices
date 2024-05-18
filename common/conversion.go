package common

import "strconv"

func Itoa(i int) string {
	return strconv.Itoa(i)
}

func Atoi(s string, defVal ...int) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return defVal[0]
	}
	return i
}
