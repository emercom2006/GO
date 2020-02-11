package main

import (
	"fmt"
	"math"
)

const DollarRate float64 = 63.57

var Rub float64

func main() {
	fmt.Println("Введите, сколько хотите обменять рублей на USD")
	if _, err := fmt.Scanln(&Rub); err != nil {
		fmt.Printf("Вы ввели не число %s", err)
		return
	}
	b := Rub / DollarRate
	fmt.Println("Вы получите", math.Round(b), "USD")
	fmt.Printf("Вы получите: %.2f USD", Rub/DollarRate)
}
