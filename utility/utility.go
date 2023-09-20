package utility

import (
	"io/ioutil"
	"log"

	"github.com/fatih/color"
)

func LogInfo(text string) {
	color.Yellow(text)
}

func LogError(text string) {
	color.Red(text)
}

func LogQuestion(text string) {
	color.Blue(text)
}

func LogSuccess(text string) {
	color.Green(text)
}

func CopyFile(src string, dest string) {
	bytesRead, err := ioutil.ReadFile(src)

	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(dest, bytesRead, 0644)

	if err != nil {
		log.Fatal(err)
	}
}
