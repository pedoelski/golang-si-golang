package main

import "fmt"

func main() {
	slice := []string{"Javascript", "Python", "Golang", "Java"}
	counter := -5

	for counter < 10 {
		if counter == 0 {
			fmt.Println("Plus")
		} else {
			fmt.Println("Counter", counter)
		}
		counter++
	}

	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Different Counter", counter)
	}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	for _, value := range slice {
		fmt.Println(value)
	}
}
