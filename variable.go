package main

import "fmt"

func main() {
	var name string
	name = "Jon Doe"
	fmt.Println(name)
	name = "Doe Jon"
	fmt.Println(name)

	age := 30
	fmt.Println(age)

	var (
		firstName = "Johny Deep"
		lastName  = "Deep Jhony"
	)

	fmt.Println(firstName, lastName)
}
