package functions

import (
	"log"
	"os"
)

func ReadFile(fileName string) []byte {
	file, err := os.ReadFile(fileName)
	if err != nil {
		log.Println("There's been a mistake!")
	}

	return file
}
