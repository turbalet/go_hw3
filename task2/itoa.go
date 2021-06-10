package task2

import (
	"github.com/turbalet/hw3/task4"
	"strings"
	"fmt"
)

func Itoa(n int) string {
	var isNegative bool
	var res strings.Builder
	// var res string
	if n < 0 {
		isNegative = true
		n = 0 - n
	}
	for n > 0 {
		rem := n % 10
		n /= 10

		res.WriteRune(rune(rem))
		//res += string(rune(rem))
	}
	if isNegative {
		res.WriteRune('-')
	}
	fmt.Println(res) // hz
	//fmt.Println(res.String())
	return task4.Reverse(res.String())
}
