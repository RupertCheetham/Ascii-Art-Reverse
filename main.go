package main

import (
	"asciiReverse/functions"
)

func main() {

	referenceFileSplitByLine, exampleFileSplitByLine := functions.ProcessTxtFiles()

	functions.FindLetters(exampleFileSplitByLine)
	functions.ConvertLetters(referenceFileSplitByLine)
	functions.PrintResult()
}
