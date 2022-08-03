package main

import "fmt"

func main() {
	name1 := "john"
	name2 := "john"

	var result1 bool = name1 == name2
	var result2 bool = name1 > name2
	var result3 bool = name1 < name2
	var result4 bool = name1 != name2

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)

	var value1 = 100
	var value2 = 200

	var result5 = value1 > value2

	fmt.Println(result5)
}
