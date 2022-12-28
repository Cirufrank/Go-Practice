// Write a program that includes an array with at least 10 values in it. The program should also include a second, empty array with 10 values. Use a for loop and the range function to populate the second array with the values in the first array.

// Challenge: Update the program so that the second array contains the values in the first array, but in the reverse order. For example, if the original array is [5 6 7], the second array should be [7 6 5].

package main

import "fmt"

func main() {
	arr1 := [...]int{1,2,3,4,5,6,7,8,9,10}
	var arr2[10]int

	//Basic
	// for index := range(arr2) {
	// 	arr2[index] = arr1[index]
	// }
	//Challenege
	for index := range(arr2) {
		arr2[index] = arr1[len(arr1) - (1 + index)]
	}

	fmt.Println(arr1)
	fmt.Println(arr2)
}