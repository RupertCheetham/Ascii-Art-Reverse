package functions

// confirms that a letter ends here
func LetterBreaksHere(exampleFileSplitByLine []string, runePosition int) bool {
	for k := 1; k < 5; k++ {
		exampleRuneofLine := []rune(exampleFileSplitByLine[k])
		if (exampleRuneofLine[runePosition]) != ' ' {
			return false
		}
	}
	return true
}
