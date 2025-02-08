package main

import (
	"fmt"
)

func main() {

	//array массив
	//неудобна ФИКСИРОВАННАЯ длина
	// var numArr1 [5]int                 // массив 5 целых чисел
	// fmt.Println(numArr1)               // [0 0 0 0 0]
	// var numArr2 [0]int                 // пустой массив
	// fmt.Println(numArr2)               // []
	// var numArr3 = [3]int{1, 2, 3}      // массив 3 целых чисел
	// fmt.Println(numArr3)               // [1 2 3]
	// var numArr4 = [...]int{1, 2, 3, 4} // авто-определение размера массива
	// fmt.Println(numArr4)               // [1 2 3 4]
	// fmt.Println(len(numArr4))          // 4

	//slice срез слайс
	//удобно
	// var list []int64 // []
	// fmt.Println("")
	// fmt.Println(list)
	// fmt.Printf("len: %v", len(list))
	// fmt.Printf("\ncap: %v", cap(list))

	// list = []int64{1, 2, 3, 4, 5} // [1 2 3 4 5]
	// fmt.Println("")
	// fmt.Println(list)
	// fmt.Printf("len: %v", len(list))
	// fmt.Printf("\ncap: %v", cap(list))

	// list = make([]int64, 0, 5) // []
	// fmt.Println("")
	// fmt.Println(list)
	// fmt.Printf("len: %v", len(list))
	// fmt.Printf("\ncap: %v", cap(list))

	// list = make([]int64, 5, 5) // [0 0 0 0 0]
	// fmt.Println("")
	// fmt.Println(list)
	// fmt.Printf("len: %v", len(list))
	// fmt.Printf("\ncap: %v", cap(list))

	// var list2 = make([]string, 5, 5) // ["","","","",""]
	// fmt.Println("")
	// fmt.Println(list2)
	// fmt.Printf("len: %v", len(list2))
	// fmt.Printf("\ncap: %v", cap(list2))

	// list2 = append(list2, "SS") // ["","","","","", "SS"]
	// fmt.Println("")
	// fmt.Println(list2)
	// fmt.Printf("len: %v", len(list2))
	// fmt.Printf("\ncap: %v", cap(list2))

	//map карта хэш-таблица мапа отображение

	// var m1 map[int32]bool
	// var m2 map[string]string
	// m3 := make(map[int]int)

	// ages := map[string]int{
	// 	"Андрей":    30,
	// 	"Анастасия": 25,
	// 	//
	// }

	// age := ages["Андрей"]
	// ages["Наталья"] = 31
	// fmt.Println(ages["Михаил"])
	// ages["Михаил"]++
	// fmt.Println(ages["Михаил"])

	// fmt.Println(m1, m2, m3, ages, age)

	//struct стурура

	type Point struct {
		X int
		Y int
	}

	p := Point{
		X: 5,
		Y: 23,
	}

	p1 := Point{1, 2}

	fmt.Println(p, p1, p.X, p.Y, p1.X, p1.Y)

}
