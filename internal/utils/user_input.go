package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Sinet2000/go-eshop-console/internal/utils/logger"
)

// Confirm prompts the user for a "yes" or "no" response and returns a boolean.
// "y", "Y", "yes", "Yes" are treated as true.
// "n", "N", "no", "No" are treated as false.
func Confirm(confirmPrompt string) (bool, error) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("%s [y/n]: ", confirmPrompt)
		input, err := reader.ReadString('\n')
		if err != nil {
			return false, fmt.Errorf("failed to read input: %w", err)
		}

		input = strings.TrimSpace(strings.ToLower(input))

		switch input {
		case "y", "yes", "ja", "jeah", "jes", "yeas":
			return true, nil
		case "n", "no", "nope", "ne", "nien", "nain":
			return false, nil
		default:
			fmt.Println("Invalid input. Please enter 'y' or 'n'.")
		}
	}
}

func PromptUserForSelection() (int, error) {
	fmt.Printf("\nSelect an option: ")
	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')
	if err != nil {
		logger.PrintlnColoredText("❗ Failed to read input. Please try again. ❗", logger.RedTxtColorCode)
		return 0, err
	}

	input = strings.TrimSpace(input)

	choice, err := strconv.Atoi(input)
	if err != nil {
		logger.PrintlnColoredText("❗ Invalid input. Please enter a valid number. ❗", logger.RedTxtColorCode)
		return 0, fmt.Errorf("invalid input: %w", err)
	}

	// fmt.Println("\nPress Enter to continue...")
	// fmt.Scanln()

	return choice, nil
}
