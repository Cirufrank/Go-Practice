// Starting with the code in Listing 9, create an array that includes 10 values, where each value is an even number between 1 and 20. You must use a for loop in your solution, but you can use a different variation on a for loop if you prefer.

package main

import "fmt"

func main() {
    var numbers [10]int
    currentEvenNum := 2

    for index := range(numbers) {
		numbers[index] = currentEvenNum
        currentEvenNum += 2
    }

    fmt.Println(numbers)
}