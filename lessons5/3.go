package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {

	file, err := os.Open("csv") // Открываем нужный csv файл
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	// Считываем его как CSV
	csvReader := csv.NewReader(file)
	csvReader.Comment = '#' //  # закоменченные строки, которые не будут выводиться.
	for {
		post, err := csvReader.Read()
		if err != nil {
			log.Println(err)
			//В цикле обходим строки и выводим их в fmt.Printf
			break
		}
		fmt.Printf("Имя: %s Фамилия: %s Телефон: %s\n", post[0], post[1], post[2])
	}

}
