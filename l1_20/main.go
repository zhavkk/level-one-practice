package main

import (
	"fmt"
	"strings"
)

func reverseOrderString(s string) string {

	words := strings.Fields(s)

	left := 0
	right := len(words) - 1

	for left < right {
		words[left], words[right] = words[right], words[left]
		left++
		right--
	}
	return strings.Join(words, " ")
}

func main() {
	str := "главрыба рыба глав"

	fmt.Println(reverseOrderString(str))

}
