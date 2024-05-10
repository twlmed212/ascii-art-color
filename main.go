package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	// renaming our package

	"ascii/ascii"
	myFunc "ascii/ascii"
)



func main() {
	// Setup the flag, and check if it's correct or not,
	// and get the arguments as well
	flagOutput, Args := myFunc.GetFlags()

	// passing File Argument, and getting the data to be processed
	txt, table := myFunc.ReadFile(Args)
	// Spliting with NewLine to get First Table
	text := strings.Split(txt, "\\n")
	// Convert Our File to 2D Table
	tabl2D := myFunc.AddingData(table)

	resultToByte := []byte{}
	// Printing in Terminal is true as default settings,
	// Once User enerd a Flag, the PrintIn Function turend off,
	// And Preparing to save the result into Slice Of byte to get final File
	if flag.NFlag() > 0 {
		if ascii.OutputFlag != "" {
			ascii.PrintIn = false
		}
	}
	//
	for i, word := range text {
		if word == "" {
			if i != 0 || !myFunc.IsEmpty(text) {
				if ascii.PrintIn {
					fmt.Println()
				} else {
					resultToByte = append(resultToByte, byte('\n'))
				}
			}
		} else {
			myFunc.PrintX(word, tabl2D, &resultToByte)
		}
	}
	if !ascii.PrintIn {
		os.WriteFile(flagOutput, resultToByte, 0o777)
	}
}
