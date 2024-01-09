package functions

// searches for the spaces between letters of the text file
// once it's found an ascii letter it adds it to an array containing them
func FindLetters(inputFileSplitByLine []string) {
	letterNum := 0
	firstLineOfInput := inputFileSplitByLine[0]
	letterStart := 0
	for i := 0; i < len(firstLineOfInput); i++ {

		if []rune(firstLineOfInput)[i] == ' ' || i == len(firstLineOfInput)-1 {
			if LetterBreaksHere(inputFileSplitByLine, i) || i == len(firstLineOfInput)-1 {
				UploadLetterToArray(letterNum, inputFileSplitByLine, letterStart, i)
				letterStart = i + 1
				letterNum++
				// if TheresASpace(inputFileSplitByLine, i) {
				// 	letterStart = letterStart + 6
				// 	letterNum++
				// }
			}
		}
	}
}
