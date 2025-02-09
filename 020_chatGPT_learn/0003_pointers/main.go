package main

import (
	"fmt"
)

func swap(a, b *int) {
	*a, *b = *b, *a
}

func main() {
	fmt.Println("Введи a: ")
	var a int
	fmt.Scanln(&a)

	fmt.Println("Введи b: ")
	var b int
	fmt.Scanln(&b)

	swap(&a, &b)
	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)

}
