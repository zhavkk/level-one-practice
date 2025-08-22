package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Разработать программу, которая будет последовательно отправлять значения в канал,
// а с другой стороны канала – читать эти значения. По истечении N секунд программа должна завершаться.

// Подсказка: используйте time.After или таймер для ограничения времени работы.

// 1.5 ----------------

func main() {
	ch := make(chan int) // создаем канал

	var n int

	fmt.Scanf("%d", &n)
	wg := sync.WaitGroup{}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wg.Add(1)
	go func(ch chan<- int) {
		defer wg.Done()
		producer(ctx, ch)
	}(ch)

	wg.Add(1)
	go func(ch <-chan int) {
		defer wg.Done()
		consumer(ctx, ch)
	}(ch)

	select {
	case <-time.After(time.Second * time.Duration(n)):
		fmt.Println("завершаем программу")
		cancel()
	}

	wg.Wait()
	fmt.Println("программа завершена")

}

func producer(ctx context.Context, ch chan<- int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("продюсер остановлен")
			return
		default:
			randomNum := rand.Intn(100)
			ch <- randomNum
			time.Sleep(500 * time.Millisecond)
		}
	}

}

func consumer(ctx context.Context, ch <-chan int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("консьюмер остановлен")
			return
		case num := <-ch:
			fmt.Printf("получено число: %d\n", num)
			time.Sleep(500 * time.Millisecond)

		}
	}
}
