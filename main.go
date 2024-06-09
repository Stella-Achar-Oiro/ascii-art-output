package main

import (
	"fmt"
	"os"
	"strings"

	ascii "ascii-art-output/banner"
)

func main() {
	var outputFileName string
	var input string

	if strings.HasPrefix(os.Args[1], "--output=") {
		outputFileName = strings.TrimPrefix(os.Args[1], "--output=")
		if len(os.Args) < 3 {
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
			fmt.Println("Example: go run . --output=<fileName.txt> something standard")
			os.Exit(1)
		}
		input = os.Args[2]
	} else {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
		fmt.Println("Example: go run . --output=<fileName.txt> something standard")
	}

	input = strings.ReplaceAll(input, "\\n", "\n")
	input = ascii.HandleSpecialCases(input)

	if input == "\n" {
		fmt.Println()
		return
	} else if input == "" {
		return
	}

	Input := strings.Split(input, "\n")
	spaceCount := 0

	var formattedOutput strings.Builder
	for _, word := range Input {
		if word == "" {
			spaceCount++
			if spaceCount < len(Input) {
				formattedOutput.WriteString("\n")
			}
		} else {
			formattedOutput.WriteString(ascii.PrintBanner(word))
			formattedOutput.WriteString("\n") // Add a newline after each word
		}
	}

	// Write the output to the specified file if --output flag is provided
	if outputFileName != "" {
		err := ascii.WriteToFile(outputFileName, formattedOutput.String())
		if err != nil {
			fmt.Printf("Error writing to file: %v\n", err)
		}
	} else {
		// Print the formatted output to the console
		fmt.Print(formattedOutput.String())
	}
}
