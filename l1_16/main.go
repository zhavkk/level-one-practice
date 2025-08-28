package main

import "fmt"

func quickSort(arr []int, start int, end int) {

	if end-start+1 <= 1 {
		return
	}

	pivot := arr[end]

	left := start

	for i := start; i < end; i++ {
		if arr[i] < pivot {
			arr[i], arr[left] = arr[left], arr[i]
			left++
		}
	}

	arr[end] = arr[left]
	arr[left] = pivot

	quickSort(arr, start, left-1)
	quickSort(arr, left+1, end)

}

func main() {
	arr := []int{3, 6, 8, 10, 1, 2, 1}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
