package main

import "sync"

func main() {
	massiv := [10]int{1, 32, 123, 432, 12, 90, 54, 23, 11, 5}
	ch1 := make(chan int, len(massiv))
	ch2 := make(chan int, len(massiv))

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func(ch chan<- int) {
		defer func() {
			wg.Done()
			close(ch1)
		}()
		for _, v := range massiv {
			ch <- v
		}
	}(ch1)

	wg.Add(1)
	go func() {
		defer func() {
			wg.Done()
			close(ch2)
		}()
		for v := range ch1 {
			ch2 <- v * v
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range ch2 {
			println(v)
		}
	}()

	wg.Wait()

}
