package functions

import "strings"

func ProcessTextFiles() ([]string, []string) {

	referenceFile := ReadFile("textFiles/standard.txt")
	exampleFile := ReadFile("textFiles/example.txt")

	referenceFileSplitByLine := strings.Split(string(referenceFile), "\n")
	exampleFileSplitByLine := strings.Split(string(exampleFile), "\n")

	return referenceFileSplitByLine, exampleFileSplitByLine
}
