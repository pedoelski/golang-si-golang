package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println("Break Counter", i)

		if i == 5 {
			break
		}
	}

	for i := 0; i < 9; i++ {
		if i%2 == 0 {
			continue
		}

		fmt.Println("Continue Counter", i)
	}
}
