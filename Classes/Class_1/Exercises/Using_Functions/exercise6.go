package main

import (
	"fmt"
	"strings"
	"math"
	"strconv"
)

func main() {
	isRunning := true;
	fmt.Println("Hello there and welcome to the calculator!")
	for(isRunning) {
		var currentChoice string
		for {
			fmt.Println("Please type 'calculator' to begin calculating or 'exit' to exit the program")
			fmt.Scan(&currentChoice)
			if (isValidUserChoice(currentChoice)) {
				break
			}
			fmt.Println("Invalid choice:",currentChoice)
		}
		currentChoice = strings.ToLower(currentChoice)
		switch(currentChoice) {
		case "exit":
			isRunning = false
		case "calculator":
			runCalculator()
		}
		
	}
	fmt.Println("Bye! Thanks so much for stopping by\nand come again soon!")
}

func runCalculator() {
	var operator string
	fmt.Println("Please input the type of operation you'd like to perform")
	printCalculatorMenu()
	fmt.Scan(&operator)
	for !isValidOperator(operator) {
		fmt.Println("Invalid operator:",operator)
		fmt.Println("Please try again")
		fmt.Scan(&operator)
	}
	operator = strings.ToLower(operator)

	switch (operator) {
	case "+":
		performAddition()
	case "-":
		performSubtraction()
	case "*":
		performMultiplication()
	case "/":
		performDivision()
	case "%":
		performGetModulus()
	case "!":
		performFactorial()
	case "sqrt":
		performSqrt()
	}
}

func printCalculatorMenu() {
	fmt.Println("Type + for addition")
	fmt.Println("Type - for subtraction")
	fmt.Println("Type * for multiplication")
	fmt.Println("Type / for division")
	fmt.Println("Type % for modulous")
	fmt.Println("Type ! for factorial")
	fmt.Println("Type 'Sqrt' for squart root")
}

func performAddition() {
	num1 := getValidNumber("Please enter the first integer")
	num2 := getValidNumber("Please enter the second integer")
	fmt.Println(num1,"+",num2,"=",addNums(num1, num2))
}

func performSubtraction() {
	num1 := getValidNumber("Please enter the first integer")
	num2 := getValidNumber("Please enter the second integer")
	fmt.Println(num1,"-",num2,"=",subtractNums(num1, num2))
}

func performMultiplication() {
	num1 := getValidNumber("Please enter the first integer")
	num2 := getValidNumber("Please enter the second integer")
	fmt.Println(num1,"*",num2,"=",multiplyNums(num1, num2))
}

func performDivision() {
	num1 := getValidNumber("Please enter the first integer")
	num2 := getValidNumber("Please enter the second integer")
	fmt.Println(num1,"/",num2,"=",divideNums(num1, num2))
}

func performGetModulus() {
	num1 := getValidNumber("Please enter the first operand")
	num2 := getValidNumber("Please enter the second operand")
	fmt.Println(num1,"%",num2,"=",findModOfNum(num1, num2))
}

func performFactorial() {
	num := getValidNumber("Please enter the number you'd like to get the factorial of")
	fmt.Println("The factorial of",num,"is",factorialOfNum(num))
}

func performSqrt() {
	num := getValidNumber("Please enter the number you'd like to get the square root of")
	fmt.Println("The square root of",num,"is",squartRootOfNum(num))
}

func addNums(num1 int, num2 int) int {
	return num1 + num2
}

func divideNums(num1 int, num2 int) float64 {
	return math.Round(float64(num1)/float64(num2) * 100)/100
}
func multiplyNums(num1 int, num2 int) int {
	return num1 * num2
}

func subtractNums(num1 int, num2 int) int {
	return num1 - num2
}

func findModOfNum(num1 int, num2 int) int {
	return num1 % num2
}

func squartRootOfNum(num int) float64{
	return math.Round(math.Sqrt(float64(num)) * 100)/100
} 

func factorialOfNum(num int) int {
	if (num == 1) {
		return num
	}

	return num * factorialOfNum(num - 1)
}

func isValidUserChoice(choice string) bool {
	return strings.ToLower(choice) == "calculator" || strings.ToLower(choice) == "exit"
}

func isValidOperator(operator string) bool {
	return strings.Contains("+-*/%!sqrt", strings.ToLower(operator))
}

func getValidNumber(prompt string) int {
	var numString string
	fmt.Println(prompt)
	fmt.Scan(&numString)
	if _, err := strconv.Atoi(numString); err != nil {
		fmt.Println("Invalid integer:", numString)
		fmt.Println("Please try again")
		fmt.Scan(&numString)
	}

	num, _ := strconv.Atoi(numString)
	return num;
}