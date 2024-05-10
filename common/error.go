package common

import "log"

func Fatal(err error, s ...string) {
	if err != nil {
		if len(s) == 0 {
			log.Fatal(err)
		} else {
			log.Fatal(s)
		}
	}
}

//
func Panic(err error, s ...string) {
	if err != nil {
		if len(s) == 0 {
			panic(err)
		} else {
			panic(s)
		}
	}
}
