package ascii

// import (
// 	"fmt"
// 	"os"
// )

// func FileCheck(fileName string) (string, error) {
// 	var fileSize int64

// 	file_info, err := os.Stat("./banners/" + fileName + ".txt")
// 	if err != nil {
// 		return fileName, err
// 	}
// 	fileSize = file_info.Size()

// 	validSizes := map[string]int64{
// 		"standard":   6623,
// 		"thinkertoy": 5503,
// 		"shadow":     7463,
// 	}

// 	if validSize, ok := validSizes[fileName]; ok && fileSize != validSize {
// 		return fileName, fmt.Errorf("the Banner file has been altered")
// 	}

// 	return fileName, nil
// }

import (
	"bufio"
	"fmt"
	"os"
)

func FileCheck(fileName string) (string, error) {
	filePath := "./banners/" + fileName + ".txt"

	file, err := os.Open(filePath)
	if err != nil {
		return fileName, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineCount := 0
	for scanner.Scan() {
		lineCount++
	}

	if lineCount != 855 {
		return fileName, fmt.Errorf("incorrect number of lines in file, expected 855 but found %d", lineCount)
	}

	if err := scanner.Err(); err != nil {
		return fileName, err
	}

	return fileName, nil
}
