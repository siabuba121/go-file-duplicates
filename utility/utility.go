package utility

import (
	"github.com/fatih/color"
)

func LogInfo(text string) {
	color.Yellow(text)
}

func LogError(text string) {
	color.Red(text)
}
