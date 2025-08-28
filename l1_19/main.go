package main

import "fmt"

func reverseString(s string) string {
	runes := []rune(s)

	left := 0
	right := len(runes) - 1

	for left < right {
		runes[left], runes[right] = runes[right], runes[left]
		left++
		right--
	}

	return string(runes)
}

func main() {
	str := "главрыба"

	fmt.Println(reverseString(str))

}
