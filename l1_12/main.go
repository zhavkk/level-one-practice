package main

import "fmt"

// Имеется последовательность строк: ("cat", "cat", "dog", "cat", "tree"). Создать для неё собственное множество.

// Ожидается: получить набор уникальных слов. Для примера, множество = {"cat", "dog", "tree"}.

func main() {

	array := []string{"cat", "cat", "dog", "cat", "tree", "tree"}

	used := make(map[string]int)

	mnoj := []string{}

	for _, v := range array {
		if used[v] > 0 {
			continue
		}
		mnoj = append(mnoj, v)
		used[v]++
	}

	fmt.Println(mnoj)

}
