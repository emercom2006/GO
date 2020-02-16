package main

import "fmt"

func main() {
	var a int
	fmt.Println("Напишите число:")

	if _, err := fmt.Scanln(&a); err !=nil {fmt.Printf("Вы ввели не число %s", err)
		return}

	if a % 3 == 0 {
		fmt.Println(a, "Делится без остатка")
	} else {
		fmt.Println(a, "Не делится без остатка")
	}

}
