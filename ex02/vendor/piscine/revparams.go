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

func RevParams() {
	var argv []string = os.Args[1:]
	var rev string
	for _, v := range argv {
		rev = (v + "\n") + rev
	}
	Print(rev)
}
