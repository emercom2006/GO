package main

func main() {
	naturals := make(chan int, 1)
	squares := make(chan int, 1)

	// генерация
	go func() {
		for x := 0; x <= 10; x++ {
			naturals <- x
		}
	}()

	// возведение в квадрат
	go func() {
		for {
			x, ok := <-naturals
			if !ok {
				break // канал закрыт и пуст
			}
			squares <- x * x
		}
	}()
}
