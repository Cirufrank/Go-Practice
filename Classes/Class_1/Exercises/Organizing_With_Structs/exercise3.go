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
	zip int
}

func (addr address) GetFullName() string {
	return addr.firstName + " " + string(addr.middleName[0]) + ". " + addr.lastName
}

func main() {
	address1 := address {name{"Roro", "Last Name", "Middle Name"}, "123 Test Street", "123 Test St 2", "Tst City", "Test State", 12345}
	fmt.Println(address1.GetFullName())
	fmt.Println(address1.addressLine1)
	fmt.Println(address1.city + ",", address1.state, address1.zip)
}