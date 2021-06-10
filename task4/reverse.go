package task4

import (
	"strings"
)

func Reverse(s string) string {
	var res strings.Builder
	bytes := []rune(s)
	for i := len(bytes) - 1; i >= 0; i-- {
		//		fmt.Print(i, "-", bytes[i], ", ")
		res.WriteRune(bytes[i])
	}
	return res.String()
}
