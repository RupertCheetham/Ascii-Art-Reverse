package functions

import (
	"log"
	"os"
)

// Reads the text file and handles errors
func ReadFile(fileName string) []byte {
	file, err := os.ReadFile(fileName)
	if err != nil {
		log.Println("There's been a mistake reading the file in ReadFile.go!")
	}

	return file
}
