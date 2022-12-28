// Write a program that includes at least three separate arrays. The program should do the following using working examples:

// Declare an array in one statement and assign values to the array in a separate statement.
// Declare an array with a defined size and assign values to the array in the same statement.
// Declare an array without specifying the size and assign values to the array in the same statement.

// You may use whatever arrays you wish, including reusing arrays from the listings in this lesson. For each array, display the array itself and its length.

package main

import "fmt"

func main() {
	var array1[5]int;
	for i := range(array1) {
		array1[i] = i + 1
	}
	array2 := [5]int{5,4,3,7,8}
	array3 := [...]int{7,8,9,7,3}
	fmt.Println(array1)
	fmt.Println(array2)
	fmt.Println(array3)
}