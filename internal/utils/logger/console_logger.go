package logger

import "fmt"

const (
	GreenTxtColorCode = "\033[32m"
	RedTxtColorCode   = "\033[31m"
	GrayTxtColorCode  = "\033[90m"
	ResetTxtColorCode = "\033[0m"
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
