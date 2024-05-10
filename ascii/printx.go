package ascii

import (
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
	// "image/color"
)

var (
	PrintIn    bool = true
	OutputFlag string
	ColorFlag  string
	Letter     string
)

func CheckNonePrintable(s string) bool {
	for _, rn := range s {
		if !(rn >= 32 && rn <= 127) {
			return false
		}
	}
	return true
}

func PrintX(s string, tabl2D [][]string, toByte *[]byte) {
	colorName, checkColor := getColor(ColorFlag)
	if !checkColor && ColorFlag != "" {
		fmt.Println("❌ Error: color not Found!!")
	}
	index := getIndex(s)
	for i := 0; i < 8; i++ {
		for j := 0; j < len(s); j++ {
			if CheckNonePrintable(string(s[j])) {
				if PrintIn {
					if checkColor {
						for _, k := range index {
							if k == j {
								/// 		 "255;218;185",
								fmt.Print("\x1B[38;2;" + colorName + "m")
								fmt.Print(tabl2D[s[j]-32][i])
								fmt.Print("\x1B[0m")
								j++
							}
						}
					}
					if j < len(s) {
						fmt.Print(tabl2D[s[j]-32][i])
					}
				} else {
					temp := tabl2D[s[j]-32][i]
					for _, char := range temp {
						*toByte = append(*toByte, byte(char))
					}
				}
			}
		}
		if PrintIn {
			fmt.Println()
		} else {
			*toByte = append(*toByte, byte('\n'))
		}
	}
}

// checking if the slice Empty or not
func IsEmpty(tab []string) bool {
	for _, char := range tab {
		if len(char) != 0 {
			return false
		}
	}
	return true
}

// turning Data to a 2D Table
func AddingData(table []string) [][]string {
	tabl2D := make([][]string, len(table))
	for i := 0; i < len(table); i++ {
		tabl2D[i] = strings.Split(table[i], "\n")
	}
	return tabl2D
}

func getIndex(s string) (index []int) {
	if len(s) < len(Letter) {
		return
	}

	for i := 0; i < len(s)-len(Letter)+1; i++ {
		if s[i:i+len(Letter)] == Letter {
			j := 0
			for j < len(Letter) { // 2  //5
				index = append(index, i+j)
				j++
			}
		}
	}
	return
}

// Reading the Ascii character file got
func ReadFile(args []string) (string, []string) {
	text := ""
	fileName := "standard"
	if ColorFlag != "" {
		PrintIn = true
		Letter = args[0]
		args = append(args[:0], args[1:]...)
	}
	if len(args) == 0 {
		text = Letter
	} else {
		text = args[0]
	}
	// fmt.Println(args)
	if len(args) > 1 {
		fileName = args[1]
	}
	input, err := os.ReadFile(string(strings.ToLower(fileName) + ".txt"))
	if err != nil {
		fmt.Println("File Not Found")
		os.Exit(1)
	}
	var tab []string
	pattern := regexp.MustCompile(`\r\n`)
	v := pattern.ReplaceAllString(string(input), "\n")
	tab = strings.Split(string(v[1:]), "\n\n")
	return text, tab
}

func GetFlags() (string, []string) {
	flag.Usage = func() {
		// Color Error
		if OutputFlag == "" {
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\nEX: go run . --output=<fileName.txt> something standard")
		} else {
			fmt.Println("Usage: go run . [OPTION] [STRING]\nEX: go run . --color=<color> <letters to be colored> \"something\"")
		}
		// Output Error
		os.Exit(0)
	}
	flag.CommandLine.SetOutput(io.Discard)
	flag.StringVar(&OutputFlag, "output", "", "To save the result into a file Ex : -output=result.txt")
	flag.StringVar(&ColorFlag, "color", "", "To Color Text EX: --color=red 'H' 'Hello' ")
	flag.Parse()

	if flag.NFlag() == 2 {
		fmt.Println("🚫 You can't use --output Flag with --color flag, Cause we can't color the text File, \n💡 Please use one flag instead")
		os.Exit(1)
		// ascii.PrintIn = false
	}
	Args := flag.Args()
	// check if the Flag is a text file or not
	FilePattern := regexp.MustCompile(`\w+\.txt$`)
	isValid := FilePattern.Match([]byte(OutputFlag))
	if !isValid && OutputFlag != "" {
		fmt.Println("Your Output file is not Valid, use Text File instead, EX : --output=result.txt")
		os.Exit(1)
	}
	if len(Args) == 0 {
		fmt.Println("Please Add more argumments!!!")
		os.Exit(1)
	}
	return OutputFlag, Args
}
