package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "JavaScript"
	names[1] = "Python"
	names[2] = "Java"
	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int{
		90, 95, 60,
	}
	fmt.Println(values)
	fmt.Println(len(values))
	fmt.Println(len(names))
}
