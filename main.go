package main

import (
	ascii "ascii-art-output/banner"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
		fmt.Println("Example: go run . --output=<fileName.txt> something standard")
		os.Exit(1)
	}

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
		input = os.Args[1]
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

	// Print the formatted output to the console
	fmt.Print(formattedOutput.String())

	// Write the output to the specified file if --output flag is provided
	if outputFileName != "" {
		err := ascii.WriteToFile(outputFileName, formattedOutput.String())
		if err != nil {
			fmt.Printf("Error writing to file: %v\n", err)
		}
	}
}
