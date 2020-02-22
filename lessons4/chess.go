package main

import "fmt"

func main() {

	motion1 := []int{-1, -2}
	motion2 := []int{1, -2}
	motion3 := []int{2, -1}
	motion4 := []int{2, 1}
	motion5 := []int{1, 2}
	motion6 := []int{-1, 2}
	motion7 := []int{-2, 1}
	motion8 := []int{-2, -1}

	var mcx int
	var mcy int
	var x int
	var y int

	fmt.Println("Укажите точку нахождения коня по оси Y от 1 до 10")
	fmt.Scanln(&y)
	fmt.Println("Укажите точку нахождения коня по оси Х от 1 до 10")
	fmt.Scanln(&x)

	mcx = x + motion1[0]
	mcy = y + motion1[1]
	if mcx > 0 && mcx < 11 && mcy > 0 && mcy < 11 {
		fmt.Printf("Фигура сможет занять точку с координатами по оси X: %v и по оси Y: %v\n", mcx, mcy)
	}

	mcx = x + motion2[0]
	mcy = y + motion2[1]
	if mcx > 0 && mcx < 11 && mcy > 0 && mcy < 11 {
		fmt.Printf("Фигура сможет занять точку с координатами по оси X: %v и по оси Y: %v\n", mcx, mcy)
	}

	mcx = x + motion3[0]
	mcy = y + motion3[1]
	if mcx > 0 && mcx < 11 && mcy > 0 && mcy < 11 {
		fmt.Printf("Фигура сможет занять точку с координатами по оси X: %v и по оси Y: %v\n", mcx, mcy)
	}

	mcx = x + motion4[0]
	mcy = y + motion4[1]
	if mcx > 0 && mcx < 11 && mcy > 0 && mcy < 11 {
		fmt.Printf("Фигура сможет занять точку с координатами по оси X: %v и по оси Y: %v\n", mcx, mcy)
	}

	mcx = x + motion5[0]
	mcy = y + motion5[1]
	if mcx > 0 && mcx < 11 && mcy > 0 && mcy < 11 {
		fmt.Printf("Фигура сможет занять точку с координатами по оси X: %v и по оси Y: %v\n", mcx, mcy)
	}

	mcx = x + motion6[0]
	mcy = y + motion6[1]
	if mcx > 0 && mcx < 11 && mcy > 0 && mcy < 11 {
		fmt.Printf("Фигура сможет занять точку с координатами по оси X: %v и по оси Y: %v\n", mcx, mcy)
	}

	mcx = x + motion7[0]
	mcy = y + motion7[1]
	if mcx > 0 && mcx < 11 && mcy > 0 && mcy < 11 {
		fmt.Printf("Фигура сможет занять точку с координатами по оси X: %v и по оси Y: %v\n", mcx, mcy)
	}

	mcx = x + motion8[0]
	mcy = y + motion8[1]
	if mcx > 0 && mcx < 11 && mcy > 0 && mcy < 11 {
		fmt.Printf("Фигура сможет занять точку с координатами по оси X: %v и по оси Y: %v\n", mcx, mcy)
	}
}
