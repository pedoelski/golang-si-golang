package main

import "fmt"

func sayHello() {
	fmt.Println("Hello World")
}

func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hallo", firstName, lastName)
}

func sum(value1 int, value2 int) int {
	return value1 + value2
}

func main() {
	sayHello()
	sayHello()
	sayHello()
	sayHelloTo("Jon", "Doe")
	fmt.Println(sum(1, 3))
}
