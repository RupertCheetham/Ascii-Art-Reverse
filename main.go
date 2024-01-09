package main

import (
	"asciiReverse/functions"
)

func main() {

	referenceFileSplitByLine, inputFileSplitByLine := functions.ProcessTextFiles()

	functions.FindLetters(inputFileSplitByLine)
	functions.ConvertLetters(referenceFileSplitByLine)
	functions.PrintResult()

}
