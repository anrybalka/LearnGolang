package main

import (
	"fmt"
	"math/rand"
	"sync" // Импортируем sync для управления горутинами
	"time"
)

func printGoRoutine(id int, seconds int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("ГоРутина %d начала работу\n", id)
	time.Sleep(time.Duration(seconds) * time.Second)
	fmt.Printf("ГоРутина %d завершила работу\n", id)
}

func randInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go printGoRoutine(i, randInt(5, 10), &wg)
	}
	wg.Wait()
	fmt.Println("Все горутины завершили работу!")
}
