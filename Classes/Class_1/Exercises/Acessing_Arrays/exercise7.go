// Create a program that includes an array with at least six integers. Duplicate the array and update it so that every other value in the new array is updated to match the value immediately ahead of it.

// For example, if the original array is [7 5 10 15 30 50], the final version of the new array is [7 7 10 10 30 30]. Do not make any changes to the original array.

package main

import "fmt"

func main() {
	arr1 := [...]int{1,2,3,4,5,6,7,8}
	arr2 := arr1

	for index := 0; index < len(arr2) - 1; index += 2 {
		arr2[index] = arr2[index + 1]
	}

	fmt.Println(arr1)
	fmt.Println(arr2)
}