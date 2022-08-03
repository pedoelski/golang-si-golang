package main

import "fmt"

func main() {
	result := 10 + 10
	fmt.Println(result)

	a := 10
	b := 10
	c := a * b

	fmt.Println(c)

	i := 10
	i += 10

	fmt.Println(i)
	i++
	fmt.Println(i)

	negative := -100
	positive := +100

	fmt.Println(positive)
	fmt.Println(negative)
}
