package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("Your are blocked")
	} else {
		fmt.Println("Welcome")
	}
}

func main() {
	fun := func(name string) bool {
		return name == "dog"
	}

	registerUser("Jon Doe", fun)
	registerUser("dog", func(name string) bool {
		return name == "dog"
	})
}
