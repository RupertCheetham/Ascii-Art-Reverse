package functions

var Result []rune

// Adds the rune value to Result, so that we can convert it into a string
func AddRuneValueToResult(referenceFileLine int) {

	// each ASCII character in textFiles/standard.txt takes up 9 lines,
	// also it starts with ASCII 32 - a space
	asciiCode := ((referenceFileLine + 1) / 9) + 32
	Result = append(Result, rune(asciiCode))

}
