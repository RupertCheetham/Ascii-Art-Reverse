package functions

// checks to see if the current reference file lines are a match with lines 2-6 of current letter
func FoundALetter(referenceFileSplitByLine []string, asciiLetter []string, i int) bool {

	asciiLetterLine := 1

	for j := i + 1; j <= i+5; j++ {
		if referenceFileSplitByLine[j] != asciiLetter[asciiLetterLine] {
			return false
		} else {
			asciiLetterLine++
		}
	}
	// os.Exit(0)
	return true
}
