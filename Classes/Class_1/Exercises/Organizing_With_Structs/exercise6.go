// Consider a different scenario where it makes sense to nest a struct in another struct. Using that scenario, create a struct that contains at least one nested struct. Using the structs, write a program that creates and displays at least two different variables using the structs.

// Examples include:

// A bank customer with multiple accounts
// A student taking multiple courses
// A course with multiple students
package main

import "fmt"

type account struct {
	accountNumber int 
	amount float64
}

type customer struct {
	firstName string
	lastName string
	account
}

func main() {
	customer1 := customer{"First Name 1", "Last Name 1", account{12343, 202.20}}
	customer2 := customer{"First Name 2", "Last Name 2", account{2345, 50000000.55}}
	fmt.Println(customer1)
	fmt.Println(customer2)
}