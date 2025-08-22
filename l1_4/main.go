package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

var wg sync.WaitGroup

func main() {
	ctx, cancel := context.WithCancel(context.Background()) // создаем контекст с возможностью отмены
	defer cancel()                                          // отменяем контекст при выходе из функции main

	ch := make(chan int)

	var n int // количество воркеров

	fmt.Scanf("%d", &n)

	wg.Add(n) // добавляем количество воркеров в WaitGroup
	for i := 0; i < n; i++ {
		go worker(ctx, i, ch)
	}

	stopCh := make(chan os.Signal, 1)
	signal.Notify(stopCh, syscall.SIGINT)

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("контекст отменен, прерываю генерацию в канал")
				return
			default:
				randomNum := rand.Intn(100)
				ch <- randomNum
				// добавляем небольшую задержку для лучшего восприятия вывода
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	<-stopCh // блокируемся до тех пор, пока не получим сигнал CTRL + C
	cancel() // отменяем контекст чтобы корректно обработать завершение воркеров

	wg.Wait() // дожидаемся завершения всех воркеров, иначе возможна утечка горутин, т.к главная горутина может завершитсья раньше

	fmt.Println("Все воркеры завершены успешно!")

}

func worker(ctx context.Context, id int, ch chan int) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			// завершаем горутину при отмене контекста
			fmt.Printf("Worker %d exiting due to context cancellation\n", id)

			return
		case num := <-ch:
			fmt.Printf("Worker %d processed number: %d\n", id, num)
		}
	}

}
