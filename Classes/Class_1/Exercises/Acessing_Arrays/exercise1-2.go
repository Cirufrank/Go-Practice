// Create a program with two different arrays. Create the arrays, add the values, and print both arrays:

// The first array should include at least 10 different integers.
// The second array should include at least five different string values.

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var arr1 [10]int
	var arr2 [5]string
	for index := range(arr1) {
		arr1[index] = index + 1
	}
	for index :=range(arr2) {
		arr2[index] = "String " + strconv.Itoa(index)
	}
	fmt.Println(arr1)
	fmt.Println(arr2)
}