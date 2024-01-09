package functions

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func ProcessTxtFiles() ([]string, []string) {

	var fileName string

	flag.StringVar(&fileName, "reverse", "", "Specify the filename")
	flag.Parse()

	if fileName == "" {
		fmt.Println("Error: Please provide a filename using the --reverse flag.")
		os.Exit(1)
	}

	referenceFile := ReadFile("textFiles/standard.txt")
	exampleFile := ReadFile("textFiles/" + fileName)

	referenceFileSplitByLine := strings.Split(string(referenceFile), "\n")
	exampleFileSplitByLine := strings.Split(string(exampleFile), "\n")

	return referenceFileSplitByLine, exampleFileSplitByLine
}
