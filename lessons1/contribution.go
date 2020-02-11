package main

import (
	"fmt"
)

func main() {
	var contribution, percent, sum float64

	fmt.Print("Введите сумму вклада:")
	if _, err := fmt.Scanln(&contribution); err != nil {
		fmt.Printf("Вы ввели не число %s", err)
		return
	}
	fmt.Print("Введите годовой процент:")
	if _, err := fmt.Scanln(&percent); err != nil {
		fmt.Printf("Вы ввели не число %s", err)
		return
	}
	for i := 1; i <= 5; i++ {
		sum = contribution * (percent / 100)
		contribution = contribution + sum
		fmt.Println("Сумма в % за", i, "год равна:", sum, "и общая сумма вклада составит:", contribution)
	}
}
