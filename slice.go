package main

import "fmt"

func main() {
	var months = [...]string{
		"January",
		"February",
		"March",
		"April",
		"Mei",
		"June",
		"July",
		"Augusts",
		"September",
		"October",
		"November",
		"December",
	}

	var slice1 = months[4:7]

	fmt.Println(slice1)
	slice1[0] = "May"
	fmt.Println(months)
	days := [...]string{
		"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday",
	}

	daySlice1 := days[5:]
	daySlice1[0] = "New Saturday"
	daySlice1[1] = "New Sunday"

	fmt.Println(days)

	daySlice2 := append(daySlice1, "Free Day")
	daySlice2[0] = "Ups"
	fmt.Println(daySlice2)
	fmt.Println(days)

	newSlice := make([]string, 2, 5)
	newSlice[0] = "Jon"
	newSlice[1] = "Doe"
	fmt.Println(newSlice)
	fmt.Println(cap(newSlice))
	fmt.Println(len(newSlice))

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	thisArray := [5]int{1, 2, 3, 4, 5}
	thisSlice := []int{1, 3, 4, 5, 5}

	fmt.Println(thisSlice)
	fmt.Println(thisArray)
}
