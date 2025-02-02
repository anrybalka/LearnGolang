package main

import "fmt"

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func main() {
	//простые string, integer, float, bool

	//[u]int[n]

	//type rune = int32 //символы unicode
	//type byte = uint8 //байты

	//float32 ~6 знаков // быстро копятся ошибки
	//float64 ~15 знаков

	//bool
	var isBool bool = true
	var numBool int = boolToInt(isBool)
	fmt.Println(numBool)

	//string
	var str1 = "\tHello\n"
	fmt.Println(str1)

	var str2 = "😂🤣😍" //unicode ok
	fmt.Println(str2)

	var str3 byte = 'c'
	fmt.Println(str3) // 99

	var str4 rune = '💋'
	fmt.Println(str4) // 128139

	var str5 = "Hello"
	fmt.Println(str5[0]) // 72
	fmt.Println(str5[1]) // 101

	var str6 = "строка"    // 6 символов
	fmt.Println(len(str6)) // 12 кол-во байтов

	var str7 = "string"    // 6 символов
	fmt.Println(len(str7)) // 6 кол-во байтов

	//str[from 0 : to len(str)]
	var str8 = "0123456789"
	var str9 = str8[:4]   // 0123
	fmt.Println(str9)     // 0123
	var str10 = str8[4:]  // 456789
	fmt.Println(str10)    // 456789
	var str11 = str8[3:6] // 345
	fmt.Println(str11)    // 345

	var str12 = "АБВГДЕЁЖЗИ"
	var str13 = str12[:4]  // АБ
	fmt.Println(str13)     // АБ
	var str14 = str12[4:]  // ВГДЕЁЖЗИ
	fmt.Println(str14)     // ВГДЕЁЖЗИ
	var str15 = str12[3:6] // �В
	fmt.Println(str15)     // �В

	var str16 = "Hello"
	fmt.Println(str16 + " World!") //Hello World!
	//str16[0] = "h" // ошибка

	//составные array, map, struct
}
