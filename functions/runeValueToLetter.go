package functions

var Result []rune

// Adds the rune value to Result, so that we can convert it into a string
func AddRuneValueToResult(referenceFileLine int) {

	asciiCode := ((referenceFileLine + 1) / 9) + 32
	Result = append(Result, rune(asciiCode))

}
