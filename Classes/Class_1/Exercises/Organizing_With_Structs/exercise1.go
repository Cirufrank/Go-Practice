package main

import "fmt"

type address struct {
	firstName string
	lastName string
	addressLine1 string
	addressLine2 string
	city string
	state string
	country string
	zip int
}

func main() {
	address1 := address{"Roro", "LastName", "1234 Test St", "1234 Test St 2", "Test City", "Test State", "Test Country", 23453}
	fmt.Println(address1.firstName,address1.lastName)
	fmt.Println(address1.addressLine1)
	fmt.Println(address1.city + ",",address1.state, address1.zip)
}