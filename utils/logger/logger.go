package logger

import (
	"fmt"
	"log"
	"os"
)

func LogError(err error) {
	file, fileErr := os.OpenFile("error.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if fileErr != nil {
		log.Fatalf("Failed to open error log file: %v", fileErr)
	}
	defer file.Close()

	// Print error to the file
	_, writeErr := fmt.Fprintf(file, "ERROR: %v\n", err)
	if writeErr != nil {
		fmt.Printf("Failed to write to error log file: %v\n", writeErr)
	}
}
