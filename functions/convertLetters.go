package functions

// converts each letter in the example file
func ConvertLetters(referenceFileSplitByLine []string) {
	for asciiLetterNum := 0; asciiLetterNum < len(ExampleFileSplitByLetterByLine); asciiLetterNum++ {
		for referenceFileLine := 0; referenceFileLine < len(referenceFileSplitByLine); referenceFileLine++ {
			if referenceFileSplitByLine[referenceFileLine] == ExampleFileSplitByLetterByLine[asciiLetterNum][0] {
				if FoundALetter(referenceFileSplitByLine, ExampleFileSplitByLetterByLine[asciiLetterNum], referenceFileLine) {
					AddRuneValueToResult(referenceFileLine)
				}
			}
		}
	}
}
