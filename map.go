package main

import "fmt"

func main() {
	person := map[string]string{
		"firstName": "Jon",
		"lastName":  "Doe",
	}

	person["title"] = "Influenced"
	person["job"] = "Programmer"

	fmt.Println(person)
	fmt.Println(len(person))
	fmt.Println(person["firstName"])
	fmt.Println(person["lastName"])
	fmt.Println(person["title"])
	fmt.Println(person["job"])
	delete(person, "job")
	fmt.Println(person)
	fmt.Println(len(person))
	fmt.Println(person["firstName"])
	fmt.Println(person["lastName"])
	fmt.Println(person["title"])
}
