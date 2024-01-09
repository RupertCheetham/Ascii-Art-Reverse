package functions

// searches for the spaces between letters of the text file
// once it's found an ascii letter it adds it to an array containing them
func FindLetters(inputFileSplitByLine []string) {
	letterNum := 0
	letterStart := 0
	for i := 0; i < len(inputFileSplitByLine[0]); i++ {

		if []rune(inputFileSplitByLine[0])[i] == ' ' || i == len(inputFileSplitByLine[0])-1 {
			if LetterBreaksHere(inputFileSplitByLine, i) || i == len(inputFileSplitByLine[0])-1 {
				UploadLetterToArray(letterNum, inputFileSplitByLine, letterStart, i)
				letterStart = i + 1
				letterNum++
				if i+5 < len(inputFileSplitByLine[0])-1 && TheresASpace(inputFileSplitByLine, i) {
					i = i + 5
					letterNum++
				}

			}
		}
	}

}
