package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
)

// 1 задание перетекает во 2

type Interface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

type customSort struct {
	t    []*PhoneBook
	less func(x, y *PhoneBook) bool
}

func (x customSort) Len() int           { return len(x.t) }
func (x customSort) Less(i, j int) bool { return x.less(x.t[i], x.t[j]) }
func (x customSort) Swap(i, j int)      { x.t[i], x.t[j] = x.t[j], x.t[i] }

// Создаем структуру телефонной книги
type PhoneBook struct {
	Name      string
	Phone     int
	Vkontakte string
	Instagram string
}

// Создаем слайс  с даными для заполнения структуры PhoneBook
var guide = []*PhoneBook{
	{"Саша", 89313481111, "vk.com/id1232", "instargam.com/vdfd"},
	{"Маша", 89313482222, "vk.com/id144", "instargam.com/veefew"},
	{"Коля", 89313483333, "vk.com/id153", "instargam.com/frwff"},
	{"Вася", 89313484444, "vk.com/id43", "instargam.com/frsd"},
}

// Создаем функцию вывода телефонного сравочника
func printPhoneBook(guide []*PhoneBook) {
	const format = "%v\t%v\t%v\t%v\t\n" // формат вывода данных заносим в константу
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Имя", "Телефон", "Вконтакте", "Инстаграм")
	fmt.Fprintf(tw, format, "---", "-------", "---------", "---------")
	for _, t := range guide {
		fmt.Fprintf(tw, format, t.Name, t.Phone, t.Vkontakte, t.Instagram)
	}
	tw.Flush() // Вычисление размеров столбцов и вывод таблицы
}

func main() {

	sort.Sort(customSort{guide, func(x, y *PhoneBook) bool {
		if x.Name != y.Name {
			return x.Name < y.Name // Сортировка по имени от А до Я
		}
		return false
	}})

	fmt.Println("\nТелефоный справочник:")

	printPhoneBook(guide)
}
