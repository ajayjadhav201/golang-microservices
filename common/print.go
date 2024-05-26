package common

import (
	"errors"
	"fmt"
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
