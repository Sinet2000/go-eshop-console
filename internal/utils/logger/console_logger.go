package logger

import "fmt"

const (
	SuccessColor = "\033[32m"
	ErrorColor   = "\033[31m"
	WarningColor = "\033[33m"
	GrayColor    = "\033[90m"
	ResetColor   = "\033[0m"
)

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func PrintlnColoredText(text string, colorCode string) {

	// Print the text with the desired color
	fmt.Print(colorCode)
	fmt.Println(text)
	fmt.Print("\033[0m") // Reset color to default after printing
}

func PrintColoredText(text string, colorCode string) {

	// Print the text with the desired color
	fmt.Print(colorCode)
	fmt.Print(text)
	fmt.Print("\033[0m") // Reset color to default after printing
}
