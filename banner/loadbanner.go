package ascii

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func LoadBanner(input string) map[rune]string {
    var height int
    banner := make(map[rune]string)
    currentChar := rune(32)
    charLine := []string{}
    filePath := "./banners/" + input + ".txt"

    file, err := os.Open(filePath)
    if err != nil {
        fmt.Println("Error opening file:", err)
        os.Exit(1)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    scanner.Scan()
    for scanner.Scan() {
        line := scanner.Text()

        if height == 8 {
            banner[currentChar] = strings.Join(charLine, "\n")
            currentChar++
            height = 0
            charLine = []string{}
        } else {
            charLine = append(charLine, line)
            height++
        }
    }
    if height > 0 {
        banner[currentChar] = strings.Join(charLine, "\n")
    }
    if err := scanner.Err(); err != nil {
        fmt.Println("Error scanning file:", err)
        os.Exit(1)
    }
    return banner
}
