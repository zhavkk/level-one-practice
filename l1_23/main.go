package main

import "fmt"

func deleteIthElement(s []int, i int) []int {
	if i < 0 || i >= len(s) {
		return s
	}

	copy(s[i:], s[i+1:])
	s[len(s)-1] = 0 // обнуляем для полной безопасности т.к если если тип будет указателем или интерфейсом то будет утечка т.к gc не сможет освободить память
	return s[:len(s)-1]
}

// получилось O(n) по времени и O(1) по памяти, можно O(1) по времени свапать последний элемент на место удаляемого
func main() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println("len:", len(arr), "cap", cap(arr), arr)
	arr = deleteIthElement(arr, 4)
	fmt.Println("len:", len(arr), "cap", cap(arr), arr)

}
