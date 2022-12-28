// Create a program that iterates through the values of an array and replaces the original values with the same values in the reverse order. For example, if the original array is [5 6 7], the updated array (with the same name) should be [7 6 5].

package main

import (
	"fmt"
	"math"
)

func main() {
	arr := [...]int {1,2,3,4,5}
	totalLoops := math.Trunc(float64(len(arr) / 2))
	index := 0;

	fmt.Println(arr)
	for totalLoops > 0 {
		temp := arr[index]
		arr[index] = arr[len(arr) - (1 + index)]
		arr[len(arr) - (1 + index)] = temp
		index++
		totalLoops--
	}
	fmt.Println(arr)
}