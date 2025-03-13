package main

import "fmt"

func sumPart(slice []int, ch chan int) {
	total := 0
	for _, v := range slice {
		total += v
	}
	ch <- total
}

func main_() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8}
	ch := make(chan int)
	mid := len(numbers) / 2

	go sumPart(numbers[:mid], ch) // Первая половина
	go sumPart(numbers[mid:], ch) // Вторая половина

	part1 := <-ch
	part2 := <-ch
	total := part1 + part2
	fmt.Println("Общая сумма:", total) // Общая сумма: 36
}
