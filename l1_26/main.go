package main

import (
	"fmt"
	"unicode"
)

func uniqueChars(s string) bool {
	mp := make(map[rune]struct{}) // struct{} занимает 0 байт, выгоднее чем bool (1 байт) или int (8 байт)

	for _, ch := range s {
		key := unicode.ToLower((ch))
		if _, exists := mp[key]; exists {
			return false
		}

		mp[key] = struct{}{}
	}
	return true
}

func main() {
	str := "abrAc"
	fmt.Println(uniqueChars(str))
}
