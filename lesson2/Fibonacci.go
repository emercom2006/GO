package main

import (
	"fmt"
	"math/big"
)

func fibon()  {
	// формулу вывел, как в big.Int запихать не разобрался
	var f, fibo int
	fiboa := 0
	fibob := 0

	for f=0; f <100; f++{
		if f == 0 {
			fibo = fiboa + fibob
			fibob = fiboa +1
			fiboa = fibo
			fmt.Println(fibo)
		} else {fibo = fiboa + fibob
			fibob = fiboa
			fiboa = fibo
			fmt.Println(fibo)
		}
	}
}


// Нашел в инете и зациклил вывод первых 100 чисел, но особо не разобрался, все значения функций так и не понял.

func fibonacci(n int) *big.Int {
	if n < 0 {
		panic("Negative arguments not implemented")
	}
	fst, _ := fib(n)
	return fst
}

// (Private) Returns the tuple (F(n), F(n+1)).
func fib(n int) (*big.Int, *big.Int) {
	if n == 0 {
		return big.NewInt(0), big.NewInt(1)
	}
	a, b := fib(n / 2)
	c := Mul(a, Sub(Mul(b, big.NewInt(2)), a))
	d := Add(Mul(a, a), Mul(b, b))
	if n%2 == 0 {
		return c, d
	} else {
		return d, Add(c, d)
	}
}


func Mul(x, y *big.Int) *big.Int {
	return big.NewInt(0).Mul(x, y)
}
func Sub(x, y *big.Int) *big.Int {
	return big.NewInt(0).Sub(x, y)
}
func Add(x, y *big.Int) *big.Int {
	return big.NewInt(0).Add(x, y)
}


func main (){
var i int
	fibon()
	for i=0; i<100; i++{
	fmt.Println(fibonacci(i))}
}
