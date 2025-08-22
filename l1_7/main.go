package main

import (
	"fmt"
	"sync"
)

// Реализовать безопасную для конкуренции запись данных в структуру map.

// Подсказка: необходимость использования синхронизации (например, sync.Mutex или встроенная concurrent-map).

// Проверьте работу кода на гонки (util go run -race).
var mu sync.Mutex
var wg sync.WaitGroup

func main() {
	mp := make(map[string]int)
	wg.Add(100)

	commonKey := "WB"
	for i := 0; i < 100; i++ {
		go safeWrite(mp, commonKey)
	}

	wg.Wait()

	fmt.Println("waited")

}

func safeWrite(mp map[string]int, key string) {
	defer wg.Done()
	mu.Lock()         // горутина блокирует мьютекс перед записью в map,
	defer mu.Unlock() // и разблокирует его после завершения записи
	mp[key]++
}
