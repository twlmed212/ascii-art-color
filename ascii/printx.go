package ascii

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"
)

var (
	PrintIn    bool = true
	OutputFlag string
	ColorFlag  string
	Letter     string
)

type ColorApiResponse struct {
	RGB struct {
		Value string `json:"value"`
	} `json:"rgb"`
}

func getColorApi(typ string, code string) (color string, isValid bool) {
	url := "https://www.thecolorapi.com/id?" + typ + "=" + code
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Sorry Mate, Color Not Found :)")
		os.Exit(1)
	}
	defer resp.Body.Close()
	var result ColorApiResponse
	json.NewDecoder(resp.Body).Decode(&result)
	RgbPattern := regexp.MustCompile(`[rgb]*\(*(\d+)[,|;]* *(\d+)[,|;] *(\d+)\)*`)
	color = RgbPattern.ReplaceAllString(result.RGB.Value, "${1};${2};${3}")
	isValid = RgbPattern.Match([]byte(result.RGB.Value))
	return
}

func CheckNonePrintable(s string) bool {
	for _, rn := range s {
		if !(rn >= 32 && rn <= 127) {
			return false
		}
	}
	return true
}

func checkColorType(checkedColor string) (res string, isValid bool) {
	HexPattern := regexp.MustCompile(`^#*(\w{6})$`)
	checkHex := HexPattern.FindAllString(checkedColor, -1)
	if checkHex != nil {
		checkHex[0] = HexPattern.ReplaceAllString(checkHex[0], "$1")
		// here Call function to convert HEX to RGB
		return getColorApi("hex", checkHex[0])
	}
	HslPattern := regexp.MustCompile(`[hsl]*\((\d+,) *(\d+%,) *(\d+%)\)`)
	checkHsl := HslPattern.FindAllString(checkedColor, -1)
	if checkHsl != nil {
		checkHsl[0] = HslPattern.ReplaceAllString(checkHsl[0], "${1}${2}${3}")
		// here Call function to convert Hsl to RGB
		return getColorApi("hsl", checkHsl[0])
	}
	RgbPattern := regexp.MustCompile(`[rgb]*\(*(\d+[,|;]* *\d+[,|;] *\d+)\)*`)
	checkRGB := RgbPattern.FindAllString(checkedColor, -1)
	if checkRGB != nil {
		/// We need to replace (,) with (;) before return the color
		RgbPattern := regexp.MustCompile(`[rgb]*\(*(\d+)[,|;]* *(\d+)[,|;] *(\d+)\)*`)
		checkRGB[0] = RgbPattern.ReplaceAllString(checkRGB[0], "${1};${2};${3}")
		return checkRGB[0], true
	}
	return "", false
}

func PrintX(s string, tabl2D [][]string, toByte *[]byte) {
	colorName, checkColor := getColor(ColorFlag)
	if !checkColor && ColorFlag != "" {
		colorName, checkColor = checkColorType(ColorFlag)
		if !checkColor {
			fmt.Println("❌ Error: color not Found!!")
		}
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
		fmt.Println("Usage: go run . [OPTION] [STRING]\nEX: go run . --color=<color> <letters to be colored> \"something\"")
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
