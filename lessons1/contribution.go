package main

import (
	"fmt"
)

var (
	contribution float64
	percent      float64
	sum          float64
)

func main() {
	fmt.Print("Введите сумму вклада:")
	if _, err := fmt.Scanln(&contribution); err != nil {
		fmt.Print("Вы ввели не число")
		return
	}
	fmt.Print("Введите годовой процент:")
	if _, err := fmt.Scanln(&percent); err != nil {
		fmt.Print("Вы ввели не число")
		return
	}
	for i := 1; i <= 5; i++ {
		sum = contribution * (percent / 100)
		contribution = contribution + sum
		fmt.Println("Сумма в % за", i, "год равна:", sum, "и общая сумма вклада составит:", contribution)
	}
}
