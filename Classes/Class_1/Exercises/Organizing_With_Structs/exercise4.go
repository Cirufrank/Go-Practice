// Using the address struct you've created in the previous listings, create an array of addresses called AddressBook. Assign at least four addresses to the array. Print the resulting addresses in a format that would look appropriate for an envelope.

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
	var addresses [4]address
	address1 := address {name{"Roro", "Last Name", "Middle Name"}, "123 Test Street", "123 Test St 2", "Tst City", "Test State", 12345}
	address2 := address{name{"First Name", "Last Name", "Middle Name"}, "1234 Test Street", "123 Test St 2", "Test City", "Test State", 12345}
	address3 := address{name{"First Name 2", "Last Name 2", "Middle Name 2"}, "1234 Test Street", "123 Test St 2", "Test City", "Test State", 12345}
	address4 := address{name{"First Name 3", "Last Name 3", "Middle Name 3"}, "1234 Test Street", "123 Test St 2", "Test City", "Test State", 12345}
	addresses[0] = address1
	addresses[1] = address2
	addresses[2] = address3
	addresses[3] = address4
	for _, curAddress := range(addresses) {
		fmt.Println(curAddress.GetFullName())
		fmt.Println(curAddress.addressLine1)
		fmt.Println(curAddress.city + ",", curAddress.state, curAddress.zip)
	}
}