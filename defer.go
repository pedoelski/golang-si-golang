package main

import "fmt"

func logging() {
	fmt.Println("Log History Function")
}

func runApplication(value int) {
	defer logging()
	fmt.Println("Run Application")
	result := 10 / value
	fmt.Println("result ", result)
}

func main() {
	runApplication(0)
}
