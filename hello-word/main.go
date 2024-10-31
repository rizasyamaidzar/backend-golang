package main

import "fmt"

func main() {
	// Manifest type data
	var firstName string = "riza"
	var lastName string
	lastName = "Syamaidzar"
	fmt.Println("Hello ", firstName, lastName+" selamat datang")

	// Type inference
	var fullName = "riza Afif Syamaidzar"

	opsionalName := "Apippp"

	fmt.Println("haii ", fullName, " ", opsionalName)
}
