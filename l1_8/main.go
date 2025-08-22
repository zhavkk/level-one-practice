package main

import "fmt"

func main() {
	var num int64

	fmt.Scanf("%d", &num)

	var i uint
	fmt.Scanf("%d", &i)

	var bit uint

	fmt.Println("введите значение бита (0 или 1)")
	fmt.Scanf("%d", &bit)

	fmt.Printf("число до изменения: %d\n", num)
	num = setBit(num, i, bit)
	fmt.Printf("число после изменения: %d\n", num)
}

func setBit(num int64, i uint, bit uint) int64 {
	if bit == 1 {
		num = num | 1<<i // or
	} else {
		num = num &^ 1 << i // and not
	}
	return num
}
