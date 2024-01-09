package functions

// converts each letter from ASCII to a rune value, adds it to Result
func ConvertLetters(referenceFileSplitByLine []string) {
	for asciiLetterNum := 0; asciiLetterNum < len(ExampleFileSplitByLetterByLine); asciiLetterNum++ {
		for referenceFileLine := 0; referenceFileLine < len(referenceFileSplitByLine); referenceFileLine++ {
			if referenceFileSplitByLine[referenceFileLine] == ExampleFileSplitByLetterByLine[asciiLetterNum][0] {
				if FoundALetter(referenceFileSplitByLine, ExampleFileSplitByLetterByLine[asciiLetterNum], referenceFileLine) {
					AddRuneValueToResult(referenceFileLine)
					break
				}
			}
		}
	}
}
