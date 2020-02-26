package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	//Создаем флаг src в котором указываем какой файл надо скопировать
	src := flag.String("src", "csv", "Имя копируемого файла")

	//Создаем флаг dst в котором указываем куда и с каким именем копируем новый файл
	dst := flag.String("dst", "csv1", "Имя скопированого файла")

	//Парсим значения флагов
	flag.Parse()
	fmt.Println("src:", *src)
	fmt.Println("dst:", *dst)

	//Функция открытия исходного файла
	oFile, err := os.Open(*src)
	if err != nil {
		log.Fatal(err)
	}
	defer oFile.Close()

	//Функция создания нового файла
	cFile, err := os.Create(*dst)
	if err != nil {
		log.Fatal(err)
	}
	defer cFile.Close()

	//Функция копирования открытого файла в новый файл
	_, err = io.Copy(cFile, oFile)
	if err != nil {
		log.Fatal(err)
	}

	// Синхронизация фиксирует текущее содержимое файла на диск
	err = cFile.Sync()
	if err != nil {
		log.Fatal(err)
	}
}
