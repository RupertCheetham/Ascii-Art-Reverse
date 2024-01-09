package functions

import (
	"log"
	"os"
)

// Reads the text file and handles errors
func ReadFile(fileName string) []byte {
	file, err := os.ReadFile(fileName)
	if err != nil {
		log.Println("Error reading filename in ReadFile.go!")
		os.Exit(0)
	}

	return file
}
