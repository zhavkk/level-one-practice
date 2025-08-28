package main

import "sync"

var increment = 0

func main() {
	wg := sync.WaitGroup{} // воздаем вэйтгруппу
	mx := sync.Mutex{}     // мьютекс

	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go worker(&mx, &wg)
	}
	wg.Wait() // ждем завершения всех горутин

	// запустив с флагом -race можно убедиться, что нету дата рейсов, если убрать
	// mx.Lock() и mx.Unlock() в воркере, то будет детектиться дата рейс

	println("increment =", increment)

}

func worker(mx *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	mx.Lock()
	defer mx.Unlock()
	increment++
}
