package main

import "fmt"

func endApp() {
	message := recover()
	if message != nil {
		fmt.Println("Error Message : ", message)
	}
	fmt.Println("End Apps")
}

func runApp(error bool) {
	defer endApp()

	if error {
		panic("Error")
	}
	fmt.Println("Apps Running")
}

func main() {
	runApp(false)
	print("Hello Jon Doe")
}
