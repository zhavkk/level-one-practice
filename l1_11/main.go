package main

import "fmt"

// Реализовать пересечение двух неупорядоченных множеств (например, двух слайсов)
//  — т.е. вывести элементы, присутствующие и в первом, и во втором.

// Пример:
// A = {1,2,3}
// B = {2,3,4}
// Пересечение = {2,3}
func main() {

	sl1 := []int{1, 2, 3}
	sl2 := []int{2, 3, 4}

	ans := []int{}

	mapa1 := make(map[int]int)

	for _, v := range sl1 {
		mapa1[v]++
	}

	for _, v := range sl2 {
		if _, ok := mapa1[v]; ok {
			ans = append(ans, v)
		}
	}

	fmt.Println(ans)

}
