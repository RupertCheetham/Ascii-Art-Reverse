package functions

import (
	"log"
	"os"
)

func ErrorChecker(exampleFileSplitByLine []string) {

	checkForValidNumberOfcharacters(exampleFileSplitByLine)
	checkForValidNumberOfLines(exampleFileSplitByLine)
}
func checkForValidNumberOfcharacters(exampleFileSplitByLine []string) {
	exampleLineLength := len(exampleFileSplitByLine[0])
	for i := 1; i < 8; i++ {

		if len(exampleFileSplitByLine[i]) != exampleLineLength {
			log.Println("Error, func checkForValidNumberOfcharacters: number of characters per line in txt file does not match")
			os.Exit(0)
		}

	}
}

func checkForValidNumberOfLines(exampleFileSplitByLine []string) {
	if len(exampleFileSplitByLine) < 8 {
		log.Println("Error, func checkForValidNumberOfLines: number of lines in txt file is invalid (needs 8 lines of chars)")
		os.Exit(0)
	}
}
