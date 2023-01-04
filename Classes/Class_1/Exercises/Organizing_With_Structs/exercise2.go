package main

import "fmt"

type name struct {
	firstName string
	lastName string
	middleName string
}

type address struct {
	name
	addressLine1 string
	addressLine2 string
	city string
	state string
	country string
	zip int
}

func main() {
	address1 := address{name{"Roro", "Last Name", "Middle name"}, "1234 Test St", "1234 Test St 2", "Test City", "Test state", "Test Country", 12345}
	fmt.Println(address1.firstName,address1.middleName,address1.lastName)
	fmt.Println(address1.addressLine1)
	fmt.Println(address1.city + ",",address1.state, address1.zip)
	// address1 := new(address)
	// address1.firstName = "Test"
	// fmt.Println(*address1)
}