package functions

import (
	"os"
	"strings"
)

func ProcessTextFiles() ([]string, []string) {

	fileName := ""

	if len(os.Args) > 0 {
		fileName = os.Args[1]
	}
	referenceFile := ReadFile("textFiles/standard.txt")
	exampleFile := ReadFile("textFiles/" + fileName)

	referenceFileSplitByLine := strings.Split(string(referenceFile), "\n")
	exampleFileSplitByLine := strings.Split(string(exampleFile), "\n")

	return referenceFileSplitByLine, exampleFileSplitByLine
}
