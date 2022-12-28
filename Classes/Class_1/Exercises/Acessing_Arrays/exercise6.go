// Create a program with a two-dimensional array that is at least 4×4 (an array with at least four array values, where each nested array also includes at least four values). The completed program should display the following:

// The values of the main array, as well as the length of the main array
// The values in the second nested array, as well as the length of the second array
// The last value in the last nested array

package main

import "fmt"

func main() {
	matrix := [4][4]int {
		{1,2,3,4},
		{5,6,7,8},
		{9,10,11,12},
		{13,14,15,16},
	}
	fmt.Println("The main array:",matrix)
	fmt.Println("The main array has a total of",len(matrix),"arrays nested within it")
	fmt.Println("The second nested array is:",matrix[1])
	fmt.Println("The second nested array has a total of", len(matrix[1]),"elements")
	fmt.Println("Additionally, the last value in the last nested array is",matrix[len(matrix) -1][len(matrix[len(matrix) -1]) - 1])
}