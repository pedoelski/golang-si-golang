package main

import "fmt"

func main() {
	name := false

	if name == true {
		fmt.Println("True")
	} else if name == false {
		fmt.Println("False")
	} else {
		fmt.Println("Default")
	}

	if length := len("Jon Doe"); length > 5 {
		fmt.Println("Long Name")
	} else {
		fmt.Println("Short Name")
	}
}
