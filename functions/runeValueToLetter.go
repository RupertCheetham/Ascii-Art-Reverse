package functions

import "fmt"

var Result []rune

// Adds the rune value to Result, so that we can convert it into a string
func AddRuneValueToResult(referenceFileLine int) {
	letterValue := (((referenceFileLine + 1) - (32 * 9)) / 9) - 32
	asciiCode := 'a' + letterValue - 1
	fmt.Printf("The letter at position %d is: %c\n", letterValue, asciiCode)
	Result = append(Result, rune(asciiCode))
}
