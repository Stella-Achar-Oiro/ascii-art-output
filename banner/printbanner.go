package ascii

import (
	"fmt"
	"os"
	"strings"
)

func PrintBanner(s string) string {
	bannerOutPut := make([][]string, 8)
	filename := ""

	if len(os.Args) == 4 {
		filename = os.Args[3]
	} else {
		filename = "standard"
	}

	filePath, err := FileCheck(filename)
	if err != nil {
		fmt.Printf("Error: bannerfile %s has been altered.\n", filename)
		os.Exit(1)
	}
	banner := LoadBanner(filePath)
	for _, char := range s {
		if char < 32 || char > 126 {
			fmt.Printf("Character out of range:%q\n", char)
			os.Exit(1)
		}
		if ascii, Ok := banner[char]; Ok {

			// If the character is found, split it into lines and append to the output
			asciiLines := strings.Split(ascii, "\n")
			for i := 0; i < len(asciiLines); i++ {
				bannerOutPut[i] = append(bannerOutPut[i], asciiLines[i])
			}
		} else {
			// If the character is not found, print an error message and continue
			fmt.Printf("Character not found: %q\n", char)
			continue
		}
	}

	// Build the assembled output lines
	var builder strings.Builder
	for _, line := range bannerOutPut {
		builder.WriteString(strings.Join(line, ""))
		builder.WriteRune('\n')
	}
	return builder.String()
}
