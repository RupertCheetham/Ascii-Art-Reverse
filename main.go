package main

import (
	"asciiReverse/functions"
)

var exampleFileSplitByLetterByLine [][6]string

func main() {

	referenceFileSplitByLine, inputFileSplitByLine := functions.ProcessTextFiles()

	// for i := 0; i < len(referenceFileSplitByLine); i++ {
	// 	if referenceFileSplitByLine[i] == inputFileSplitByLine[0] {
	// 		if functions.CheckForLetter(referenceFileSplitByLine, inputFileSplitByLine, i) {
	// 			fmt.Println("We've found it at line: ", i+1)
	// 			fmt.Println("That is letter:", (((i+1)-(32*9))/9)-32)
	// 		} else {
	// 			i = i + 5
	// 		}
	// 	}
	// }

	findSpaces(inputFileSplitByLine)
}

func findSpaces(inputFileSplitByLine []string) {
	firstLineOfInput := inputFileSplitByLine[0]
	exampleRuneofLine := []rune(firstLineOfInput)
	for i := 0; i < len(firstLineOfInput)-1; i++ {
		if i > 0 {
			if exampleRuneofLine[i] == ' ' {
				if letterBreaksHere(inputFileSplitByLine, i) {
					uploadLetterToArray(inputFileSplitByLine, i)
					i = 0
				}
			}
		} else {
			break
		}

	}
}

func letterBreaksHere(exampleFileSplitByLine []string, runePosition int) bool {
	for k := 1; k < 5; k++ {
		exampleRuneofLine := []rune(exampleFileSplitByLine[k])
		if (exampleRuneofLine[runePosition]) != ' ' {
			return false
		}
	}
	return true
}

func uploadLetterToArray(inputFileSplitByLine []string, runeToCutAt int) {

	exampleFileSplitByLetterByLine = append(exampleFileSplitByLetterByLine)
}
