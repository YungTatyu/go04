package piscine

import (
	"ft"
	"os"
)

func Print(s string) {
	for _, v := range s {
		ft.PrintRune(v)
	}
}

func PrintParams() {
	var argv []string = os.Args[1:]
	for _, v := range argv {
		Print(v + "\n")
	}
}
