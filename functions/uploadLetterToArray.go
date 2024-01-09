package functions

var ExampleFileSplitByLetterByLine [][]string

// adds the ascii letter to the array ExampleFileSplitByLetterByLine
func UploadLetterToArray(letterNum int, inputFileSplitByLine []string, runeToStartAt int, runeToEndAt int) {
	var asciiLetter []string
	for i := 0; i < 6; i++ {
		asciiLetter = append(asciiLetter, string(inputFileSplitByLine[i][runeToStartAt:runeToEndAt])+" ")
	}
	ExampleFileSplitByLetterByLine = append(ExampleFileSplitByLetterByLine, asciiLetter)
}
