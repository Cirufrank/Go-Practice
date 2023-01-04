package main

import (
	"fmt"
	// "math"
	"strings"
	"bufio"
	"os"
)

type burger struct {
	bun bool
	price float64
	dressed bool
}

type drink struct {
	price float64
	drinkType string
}

type side struct {
	price float64
	sideType string
}

type combo struct {
	burger
	side
	drink
	price float64
}

func determineTrueOrFalse(yesOrNo string) bool {
	return yesOrNo == "yes"
}

func getYesOrNo(prompt string, scan *bufio.Scanner) string {
	var userChoice string;
	fmt.Println(prompt)
	for userChoice != "yes" && userChoice != "no" {
		userChoice = getUserChoice("Please input yes or no", scan)
		if (userChoice != "yes" && userChoice != "no") {
			fmt.Println("Invalid input")
		}
	}
	return userChoice
}

func contains(items []string, item string) bool {
	for _, curItem := range(items) {
		if curItem == item {
			return true
		}
	}
	return false
}

func getUserChoice(prompt string, scan *bufio.Scanner) string {
	var userChoice string
	fmt.Println(prompt)
	if (scan.Scan()) {
		userChoice = scan.Text()
		userChoice = strings.TrimSpace(strings.ToLower(userChoice))
	}
	return userChoice
}

func (burg *burger) determineBurgerPrice() {
	var total float64
	if (burg.bun) {
		total = float64(7)
	} else {
		total = float64(6)
	}
	burg.price = total
}

func user_input_burger(scan *bufio.Scanner) burger {
	burger := new(burger);
	var withBun bool
	var dressed bool
	withBun = determineTrueOrFalse((getYesOrNo("Would you like your burger to have a bun?",scan)))
	dressed = determineTrueOrFalse(getYesOrNo("Would you like your burger to be dressed?",scan))
	burger.bun = withBun
	burger.dressed = dressed
	burger.determineBurgerPrice()
	return *burger
}

func (userDrink *drink) determineDrinkPrice() {
	switch userDrink.drinkType{
	case "fanta":
		userDrink.price = float64(1.75)
	case "sprite":
		userDrink.price = float64(1.80)
	case "water", "none":
		userDrink.price = float64(0)
	}
}

func user_input_drink(scan *bufio.Scanner) drink {
	userDrink := new(drink)
	var drinkChoice string
	validDrinkChoices := []string{"fanta", "sprite", "water", "none"}
	for !contains(validDrinkChoices, drinkChoice) {
		drinkChoice = getUserChoice("What drink would you like?\nPlease choose from Fanta, Sprite, Water, or None",scan)
		if !contains(validDrinkChoices, drinkChoice) {
			fmt.Println("Invalid entry")
		}
	}
	userDrink.drinkType = drinkChoice
	userDrink.determineDrinkPrice()
	return *userDrink
}

func (userSide *side) determineSidePrice() {
	switch (userSide.sideType) {
	case "fries":
		userSide.price = float64(1)
	case "onion rings":
		userSide.price = float64(1.75)
	case "salad":
		userSide.price = float64(.50)
	case "coleslaw":
		userSide.price = float64(1.25)
	case "none":
		userSide.price = float64(0)
	}
}

func user_input_side(scan *bufio.Scanner) side {
	userSide := new(side)
	var sideChoice string
	validSideChoices := []string{"fries","onion rings","salad","coleslaw","none"}
	for !contains(validSideChoices, sideChoice) {
		sideChoice = getUserChoice("What type of side would you like?\nPlease choose from fries, onion rings, salad, coleslaw, or none", scan)
		if !contains(validSideChoices, sideChoice) {
			fmt.Println("Invalid input")
		}
	}
	userSide.sideType = sideChoice
	userSide.determineSidePrice()
	return *userSide
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	burger := user_input_burger(scanner)
	userDrink := user_input_drink(scanner)
	userSide := user_input_side(scanner)
}