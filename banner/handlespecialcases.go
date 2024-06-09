package ascii

import (
	"fmt"
	"os"
	"strings"
)

func HandleSpecialCases(s string) string {
	cases := map[string]bool{
		"\\a": true, "\\t": true, "\\b": true, "\\v": true, "\\r": true, "\\f": true,
	}
	for char := range cases {
		if strings.Contains(s, char) {
			fmt.Printf("Special case %q is not supported\n", char)
			os.Exit(1)
		}
	}
	return s
}
