package main

import "fmt"

type Filter func(string2 string) string

func sayHelloWithFilter(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

func spamFilter(name string) string {
	if name == "dog" {
		return "..."
	} else {
		return name
	}

}

func main() {
	sayHelloWithFilter("Jon", spamFilter)
}
