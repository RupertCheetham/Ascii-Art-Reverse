package functions

// checks to see if the current reference file lines are a match with lines 2-6 of current letter
func FoundALetter(referenceFileSplitByLine []string, asciiLetter []string, i int) bool {

	k := 1

	for j := i + k; j <= i+5; j++ {
		if referenceFileSplitByLine[j] == asciiLetter[k] {
			k++
			continue
		} else {
			return false
		}
	}

	return true
}
