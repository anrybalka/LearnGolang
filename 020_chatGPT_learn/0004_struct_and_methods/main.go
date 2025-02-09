package main

import (
	"fmt"
)

type Car struct {
	Brand string
	Speed int
}

func (c *Car) speedAdd(value int) {
	var newSpeed = c.Speed + value
	if newSpeed > 250 {
		newSpeed = 250
		fmt.Print("Достигнута максимальная скорость 250! ")
	}
	fmt.Printf("Меняем скорость %s с %v на %v\n", c.Brand, c.Speed, newSpeed)
	c.Speed = newSpeed
}

func (c *Car) speedStop() {
	fmt.Printf("Меняем скорость %s с %v на %v\n", c.Brand, c.Speed, 0)
	c.Speed = 0
}

func main() {
	tesla1 := Car{
		Brand: "Tesla",
		Speed: 0,
	}
	tesla1.speedAdd(60)
	tesla1.speedAdd(60)
	tesla1.speedAdd(60)
	tesla1.speedAdd(60)
	tesla1.speedAdd(60)
	tesla1.speedAdd(60)
	tesla1.speedStop()
}
