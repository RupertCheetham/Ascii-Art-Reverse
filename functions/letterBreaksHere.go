package functions

// confirms that a letter ends here by checking for a space in this position on every line
func LetterBreaksHere(exampleFileSplitByLine []string, runePosition int) bool {
	for k := 5; k > 0; k-- {
		exampleRuneofLine := []rune(exampleFileSplitByLine[k])
		if (exampleRuneofLine[runePosition]) != ' ' {
			return false
		}
	}
	return true
}
