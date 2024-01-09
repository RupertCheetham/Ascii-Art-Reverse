package functions

// checks to see if the next letter is a space and skips ahead if so;
// prevents errors relating to how I split up the ASCII letters
func TheresASpace(inputFileSplitByLine []string, x int) bool {
	spaceString := "     "
	for i := 0; i <= 5; i++ {
		if string(inputFileSplitByLine[i][x:x+5]) != spaceString {
			return false
		}
	}
	return true
}
