package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {
	PhoneBook := make(map[string][]int)

	PhoneBook["Alexandr"] = []int{89311233232}
	PhoneBook["Andrey"] = []int{89051233232}
	PhoneBook["Alexandr"] = append(PhoneBook["Alexandr"], 89991233232)

	fmt.Println(PhoneBook)

	for name, phones := range PhoneBook {
		fmt.Println("Имя:", name)
		for i, phone := range phones {
			fmt.Printf("\t Телефон №%v %v \n", i+1, phone)
		}
	}

	PhoneBookJson, err := json.Marshal(PhoneBook)
	if err != nil {
		fmt.Printf("Ошибка конвертации в json %v", err)
	}
	// Визуально проверяю json
	fmt.Printf("%s", PhoneBookJson)

	err = ioutil.WriteFile("PhoneBook.json", PhoneBookJson, 0644)
	if err != nil {
		fmt.Printf("Ошибка записи в файл: %v", err)
	}

}
