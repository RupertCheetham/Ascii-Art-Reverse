package functions

func TheresASpace(inputFileSplitByLine []string, x int) bool {
	spaceString := "     "
	for i := 0; i <= 5; i++ {
		if string(inputFileSplitByLine[i][x:x+5]) != spaceString {
			return false
		}
	}
	return true
}
