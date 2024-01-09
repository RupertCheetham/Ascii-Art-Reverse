package main

import (
	"asciiReverse/functions"
	"fmt"
)

func main() {

	referenceFileSplitByLine, inputFileSplitByLine := functions.ProcessTextFiles()

	functions.FindLetters(inputFileSplitByLine)
	functions.ConvertLetters(referenceFileSplitByLine)
	fmt.Println("result:", string(functions.Result))
}
