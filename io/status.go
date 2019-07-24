package io

import "fmt"

func Error(args ...interface{}) {
	fmt.Print(" !     ")
	fmt.Println(args...)
}

func Errorf(format string, args ...interface{}) {
	fmt.Printf(" !     "+format, args...)
}

func Warning(args ...interface{}) {
	fmt.Print("  /!\\  ")
	fmt.Println(args...)
}

func Warningf(format string, args ...interface{}) {
	fmt.Printf("  /!\\  "+format, args...)
}

func Status(args ...interface{}) {
	fmt.Print("-----> ")
	fmt.Println(args...)
}

func Statusf(format string, args ...interface{}) {
	fmt.Printf("-----> "+format, args...)
}

func Info(args ...interface{}) {
	fmt.Print("       ")
	fmt.Println(args...)
}

func Infof(format string, args ...interface{}) {
	fmt.Printf("       "+format, args...)
}
