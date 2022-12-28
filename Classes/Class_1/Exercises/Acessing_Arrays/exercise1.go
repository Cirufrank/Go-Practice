// Create a program with two different arrays. Create the arrays, add the values, and print both arrays:

// The first array should include at least 10 different integers.
// The second array should include at least five different string values.

package main

import "fmt"

func main() {
	var array1 = [...]int{1,2,3,4,5,6,7,8,9,10}
	var array2 = [...]string{"one", "two", "three", "four", "five"}
	var totalValues int = 0;
	var totalStrings string = ""

	for _, num := range(array1) {
		totalValues += num
	}
	for _, text := range(array2) {
		totalStrings += text;
	}
	fmt.Println(array1)
	fmt.Println("Sum of values:",totalValues)
	fmt.Println(array2)
	fmt.Println("Sum of strings:",totalStrings)

}