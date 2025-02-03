package main

import "fmt"

func main() {

	//array массив
	//неудобна ФИКСИРОВАННАЯ длина
	var numArr1 [5]int                 // массив 5 целых чисел
	fmt.Println(numArr1)               // [0 0 0 0 0]
	var numArr2 [0]int                 // пустой массив
	fmt.Println(numArr2)               // []
	var numArr3 = [3]int{1, 2, 3}      // массив 3 целых чисел
	fmt.Println(numArr3)               // [1 2 3]
	var numArr4 = [...]int{1, 2, 3, 4} // авто-определение размера массива
	fmt.Println(numArr4)               // [1 2 3 4]
	fmt.Println(len(numArr4))          // 4

	//slice срез слайс
	//удобно

	//map отображение

	//struct стурура
}
