package common

import "os"

func EnvString(Key string, option string) string {
	val := os.Getenv(Key)
	if val != "" {
		return val
	}
	return option
}
