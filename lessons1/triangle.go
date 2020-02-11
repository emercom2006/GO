package main

import (
	"fmt"
	"math"
)

var (
	a float64
	b float64
	c float64
	s float64
	p float64
)

func main() {
	fmt.Println("Введите, значение катета a")
	if _, err := fmt.Scanln(&a); err != nil {
		fmt.Println("Вы ввели не число")
		return
	}
	fmt.Println("Введите, значение катета b")
	if _, err := fmt.Scanln(&b); err != nil {
		fmt.Println("Вы ввели не число")
		return
	}

	c := math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))

	fmt.Printf("Гипотенуза c = %.2f\n", math.Sqrt(math.Pow(a, 2)+math.Pow(b, 2)))
	fmt.Printf("Площадь треугольника S = %.2f\n", (a*b)/2)
	fmt.Printf("Периметр треугольника P = %.2f", a+b+c)
}
