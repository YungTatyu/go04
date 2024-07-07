package piscine

import (
	"ft"
	"os"
)

func PrintProgramName() {
	var programName string = os.Args[0]
	for _, v := range programName {
		ft.PrintRune(v)
	}
}
