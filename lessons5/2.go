package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	file, err := os.Open("задание1")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	data := make([]byte, 64) // Правильнее сделать так, так как в методичке мы загружаем полностью размер файла

	/*
		Для считывания данных определяется срез из 64 байтов.
		В бесконечном цикле содержимое файла считывается в срез,
		а когда будет достигнут конец файла, то есть мы получим
		ошибку io.EOF, то произойдет выход из цикла.
	*/

	for {
		n, err := file.Read(data)
		if err == io.EOF { // если конец файла
			break // выходим из цикла
		}
		fmt.Print(string(data[:n]))
	}

}
