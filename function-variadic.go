package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}

func main() {
	total := sumAll(3, 4, 5, 1, 23, 3, 2, 3, 3, 123, 11)
	fmt.Println(total)
	slice := []int{10, 20, 30, 40}
	total = sumAll(slice...)
	fmt.Println(total)
}
