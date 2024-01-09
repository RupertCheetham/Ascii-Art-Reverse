package functions

import "fmt"

func CheckForLetter(referenceFileSplitByLine []string, inputFileSplitByLine []string, i int) bool {

	fmt.Println(referenceFileSplitByLine[i])
	k := 1
	for j := i + k; j < i+6; j++ {
		if referenceFileSplitByLine[j] == inputFileSplitByLine[k] {
			fmt.Println(referenceFileSplitByLine[j])
			k++
			continue
		} else {
			fmt.Println("nope")
			return false
		}
	}

	return true
}
