package print

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

type ColorLevel string

const (
	Error   ColorLevel = "red"
	Warning ColorLevel = "yellow"
	Info    ColorLevel = "blue"
	Success ColorLevel = "green"
	None    ColorLevel = "none"
)

func ExitOnError(str string, err error) {
	fmt.Println("")
	if str != "" {
		PrintMessage(str, Error)
	}
	if err != nil {
		fmt.Printf("Error message: %s\n", err)
		PrintMessage("Unexpected error", Error)
	}
	os.Exit(1)
}

func PrintMessage(message string, level ColorLevel) {
	switch level {
	case Error:
		color.Red(message)
	case Warning:
		color.Yellow(message)
	case Info:
		color.Blue(message)
	case Success:
		color.Green(message)
	default:
		color.White(message)
	}
}
