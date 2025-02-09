package main

import (
	"fmt"
	"math/rand" // Импортируем sync для управления горутинами
	"time"
)

func printGoRoutine(id int, seconds int, ch chan string) {
	fmt.Printf("ГоРутина %d начала работу\n", id)
	time.Sleep(time.Duration(seconds) * time.Second)
	ch <- fmt.Sprintf("Горутина %d завершилась", id)
}

func randInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

func main() {
	ch := make(chan string) // Создаём канал
	for i := 1; i <= 10; i++ {
		go printGoRoutine(i, randInt(5, 10), ch)
	}
	for i := 1; i <= 10; i++ {
		fmt.Println(<-ch)
	}

}
