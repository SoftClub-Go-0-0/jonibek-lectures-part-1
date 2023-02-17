package main

import (
	"fmt"
	"math"
)

func main() {
	var n float64
	fmt.Print("Введите целое число: ")
	fmt.Scanln(&n)
	var a, b float64
	fmt.Print("A = ")
	fmt.Scanln(&a)
	fmt.Print("B = ")
	fmt.Scanln(&b)
	h := (b-a+1)/n
	for i := a; i < b; i+=h {
		fmt.Println(h, 1-math.Sin(i))
	}
}