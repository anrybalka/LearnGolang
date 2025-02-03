package main

import "fmt"

func main() {

	x := 1
	pointX := &x
	fmt.Println(x)
	*pointX = 2
	fmt.Println(x)
}
