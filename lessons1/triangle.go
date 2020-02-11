package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b float64
	fmt.Println("Введите, значение катета a")
	if _, err := fmt.Scanln(&a); err != nil {
		fmt.Printf("Вы ввели не число %s", err)
		return
	}
	fmt.Println("Введите, значение катета b")
	if _, err := fmt.Scanln(&b); err != nil {
		fmt.Printf("Вы ввели не число %s", err)
		return
	}

	c := math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))

	fmt.Printf("Гипотенуза c = %.2f\n", math.Sqrt(math.Pow(a, 2)+math.Pow(b, 2)))
	fmt.Printf("Площадь треугольника S = %.2f\n", (a*b)/2)
	fmt.Printf("Периметр треугольника P = %.2f", a+b+c)
}
