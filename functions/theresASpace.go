package functions

import "fmt"

func TheresASpace(inputFileSplitByLine []string, x int) bool {
	spaceString := "     "
	for i := 0; i <= 5; i++ {
		fmt.Println("X", inputFileSplitByLine[i][x:x+5])
		if string(inputFileSplitByLine[i][x:x+5]) != spaceString {
			return false
		}
	}
	fmt.Println("space")
	return true
}
