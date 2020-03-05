package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {

	//Создаем флаг grep в котором указываем в каком файле искать
	grep := flag.String("grep", "csv", "В каком файле искать?")

	//Создаем флаг search в котором указываем что искать
	search := flag.String("search", "Александр", "Что искать")

	//Парсим значения флагов
	flag.Parse()
	fmt.Println("grep:", *grep)
	fmt.Println("search:", *search)

	f, err := os.Open(*grep)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	// Разбиваем на строки
	scanner := bufio.NewScanner(f)

	line := 1

	for scanner.Scan() {
		if strings.Contains(scanner.Text(), *search) {
			fmt.Printf("Ваш запрос находится в строке #%v в файле %v", line, *grep)
		}

		line++
	}

	if err := scanner.Err(); err != nil {
	}
}
