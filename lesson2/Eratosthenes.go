package main

import "fmt"

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func main() {
var i int
		a := makeRange(2, 102)
		fmt.Println(a) // выводим значения из задания 4.a

	for i = a[0]; i <= (len(a)+1); i++ {

		if a[i-2]%a[0] == 0 {
//ненужные значения откидываем
		} else {
			fmt.Printf("Число %v не кратно переменой p = %v\n", a[i-2], a[0])

			// здесь должен формироваться новый массив допустим arr со значениями в цикле a[i-2]
			// не знаю как правильно это записать arr := make([]int, a[i-2])
			// и далее в цикле прогоняться уже с условиями по формуле предыдущего цикла

			}
		}
}