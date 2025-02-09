package main

import (
	"errors"
	"fmt"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("[ERROR]: делитель = 0")
	}
	return a / b, nil
}

func main() {
	fmt.Println("Введи делимое: \n")
	var a int
	fmt.Scanln(&a)

	fmt.Println("Введи делитель: \n")
	var b int
	fmt.Scanln(&b)

	if result, err := divide(a, b); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("\nОтвет %v", result)
	}
}
