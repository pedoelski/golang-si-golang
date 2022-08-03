package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var noKtpJohny NoKTP = "123123123123123"
	var statusJohny Married = false

	fmt.Println(noKtpJohny)
	fmt.Println(statusJohny)
}
