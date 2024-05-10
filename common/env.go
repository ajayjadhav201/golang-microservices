package common

import "syscall"

func EnvString(Key string, option string) string {
	val, ok := syscall.Getenv(Key)
	if ok {
		return val
	}
	return option
}
