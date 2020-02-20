package main

import (
	"fmt"
	"math/big"
)

func main() {
	var f int
	fiboa := big.NewInt(0)
	fibob := big.NewInt(1)

	for f = 0; f < 100; f++ {
		if f == 0 {
			fiboa = fiboa.Add(fiboa, fibob)
			fiboa, fibob = fibob, fiboa
			fmt.Println(fiboa)
		} else {
			fiboa = fiboa.Add(fiboa, fibob)
			fibob, fiboa = fiboa, fibob
			fmt.Println(fiboa)
		}
	}
}
