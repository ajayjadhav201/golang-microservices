package common

import (
	"errors"
	"fmt"
	"log"
)

func Println(a ...any) {
	fmt.Println(a...)
}

func Printf(format string, a ...any) {
	fmt.Printf(format, a...)
}

func Print(a ...any) {
	fmt.Print(a...)
}

func SPrint(a ...any) string {
	return fmt.Sprint(a...)
}

func Sprintf(format string, a ...any) string {
	return fmt.Sprintf(format, a...)
}

func SPrintln(a ...any) string {
	return fmt.Sprintln(a...)
}

func Errorf(format string, a ...any) error {
	return fmt.Errorf(format, a...)
}

func Error(err string) error {
	return errors.New(err)
}

func Fatal(err error, s ...string) {
	if err != nil {
		if len(s) == 0 {
			log.Fatal(err)
		} else {
			log.Fatal(s)
		}
	}
}

func Fatalf(format string, v ...any) {
	log.Fatalf(format, v...)
}

func Panic(err error, s ...string) {
	if err != nil {
		if len(s) == 0 {
			panic(err)
		} else {
			panic(s)
		}
	}
}
