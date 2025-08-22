package main

import (
	"context"
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

// Реализовать все возможные способы остановки выполнения горутины.

// Классические подходы: выход по условию, через канал уведомления, через контекст, прекращение работы runtime.Goexit() и др.

// Продемонстрируйте каждый способ в отдельном фрагменте кода.
var wg sync.WaitGroup

func main() {

	// выход по условию
	wg.Add(1)
	go stopByCondition()

	// канал уведомления
	stopCh := make(chan struct{})
	wg.Add(1)
	go stopByChannel(stopCh)

	// контекст
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	wg.Add(1)
	go stopByContext(ctx)

	// goexit
	wg.Add(1)
	go stopByGoexit()

	// паника
	wg.Add(1)
	go stopByPanic()

	time.Sleep(time.Second) // даем время для выполнения горутин
	close(stopCh)           // закрываем канал уведомления
	cancel()                // отменяем контекст
	// ждем завершения всех горутин
	wg.Wait()
	fmt.Println("все горутины завершены")

}

func stopByCondition() {
	flag := rand.Intn(10)
	defer wg.Done()
	for i := 0; i < 11; i++ {
		if i == flag {
			fmt.Println("stop by condition at", i)
			return
		}
	}
}

func stopByChannel(ch <-chan struct{}) {
	defer wg.Done()

	for {
		select {
		case <-ch:
			fmt.Println("stopped by channel")
			return
		default:
			randomNum := rand.Intn(100)
			fmt.Println("random number ", randomNum)
			time.Sleep(250 * time.Millisecond)
		}
	}
}

func stopByContext(ctx context.Context) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("stopped by context")
			return
		default:
			randomNum := rand.Intn(100)
			fmt.Println("random number ", randomNum)
			time.Sleep(250 * time.Millisecond)
		}
	}
}

func stopByGoexit() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		if i == 5 {
			fmt.Println("stopping by Goexit at", i)
			runtime.Goexit()
		}
		fmt.Println("iteration go exit", i)
	}
}

func stopByPanic() {
	defer func() {
		wg.Done()
		if r := recover(); r != nil {
			fmt.Println("recovered from panic:", r)
			return
		}
	}()
	for i := 0; i < 10; i++ {
		if i == 5 {
			fmt.Println("stopping by panic at", i)
			panic("panic triggered")
		}
		fmt.Println("iteration by panic", i)
	}
}
