package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch := make(chan int) // создаем канал

	var n int // количество воркеров
	fmt.Scanf("%d", &n)

	for i := 0; i < n; i++ { // запускаем n воркеров
		go worker(i, ch)
	}

	// на данном этапе n воркеров блокируются на чтении из канала ch, до тех пор пока не появится запись в канал

	// в главной горутине генерируем случайные числа и отправляем их в канал ch
	for {
		randomNum := rand.Intn(100)
		ch <- randomNum
		time.Sleep(500 * time.Millisecond) // задержка для лучшего воспрития вывода
	}

	// как только данные поступят в канал, какой то из воркеров прочитает данные и запишет в stdout,
	// какой именно воркер - зависит от планировщика, после того как запишет в stdout воркер снова блокируется на чтении из канала

}

func worker(id int, ch chan int) {
	for num := range ch {
		fmt.Printf("worker %d processed number: %d\n", id, num)
	}
}
