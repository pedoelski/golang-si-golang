package main

import "fmt"

func main() {
	value := 100

	switch value {
	case 50:
		fmt.Println("Value 50")
	case 60:
		fmt.Println("Value 60")
	default:
		fmt.Println("Default")
	}

	switch length := 50; length > 5 {
	case true:
		fmt.Println("True")
	case false:
		fmt.Println("False")
	}
}
