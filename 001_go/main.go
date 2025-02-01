package main

import "fmt"

type My struct {
	a int
	b int
}

func NewMy(a,b int) *My{
	return &My{
		a: a,
		b: b,
	}
}

func main()  {
	s:=NewMy(1,2)
	fmt.
}
