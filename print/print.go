package print

import "github.com/fatih/color"

type ColorLevel string

const (
	Error   ColorLevel = "red"
	Warning ColorLevel = "yellow"
	Info    ColorLevel = "blue"
	Success ColorLevel = "green"
	None    ColorLevel = "none"
)

func Message(message string, level ColorLevel) {
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
