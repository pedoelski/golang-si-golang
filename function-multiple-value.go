package main

import "fmt"

func getFullName() (string, string, string) {
	return "Jon", "Middle", "Doe"
}

func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "Jane"
	middleName = "Middle"
	lastName = "Doe"

	return
	//return firstName, middleName, lastName
}

func main() {
	firstName, _, lastName := getFullName()

	fmt.Println(firstName)
	fmt.Println(lastName)

	a, b, c := getCompleteName()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
