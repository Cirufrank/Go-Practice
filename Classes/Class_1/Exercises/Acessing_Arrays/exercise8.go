// Create a program that performs the following tasks:

// Prompt the user to answer a series of 5–10 questions about themselves, such as their name, their birthday, where they live, their favorite hobby/sport, etc.
// Include at least 10 questions but give the user the option to quit after answering at least five questions.
// Save the answers in an array.
// Display the results to the user in a user-friendly format.
// For example, if one of the questions is “What is your name?,” the output for that response should look like “Your name is Mary.”
// After displaying all answers, prompt the user to change one or more of the answers.
// Update the array with the new answers and redisplay the results.
// Allow the user the option to exit the program at any time.

package main

import (
	"fmt"
	"bufio"
	"strings"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	questions := [...]string{"What is you favorite passtime?",
				"Who do you look up to the most",
				"What school did you attend when you were younger?",
				"What is your favorite color",
				"What hobby makes you the happiest",
				"Do you like sports?",
				"What is your favorite meal?"}
	var answers [len(questions)]string
	answerQuestions := true
	initialRound := true
	for answerQuestions {
		for index, question :=range questions {
			if (index > 4) {
				continueProgram := determineChoice("Would you like to continue",scanner)
				if (!continueProgram) {
					break;
				}
			}
			fmt.Println(question)
			if (!initialRound) {
				prevAnswer := answers[index];
				if (prevAnswer != "") {
					fmt.Println("Enter an empty space to not change your previous answer:", prevAnswer)
				}
			}
			if scanner.Scan() {
				newAnswer := scanner.Text()
				if (!initialRound) {
					if (newAnswer == "") {
						continue;
					} 
				}
				answers[index] = newAnswer
			}
		}
		fmt.Println("Summary:")
		for index, answer := range answers {
			if (answer == "") {
				break;
			}
			currentQuestion := questions[index]
			fmt.Println(currentQuestion +":", answer)
		}
		if (initialRound) {
			initialRound = false
		}
		answerQuestions = determineChoice("Would you like to change your answers?", scanner)
	}
	fmt.Println("That's all folks! Thanks for coming and sharing a bit about yourself")
}

func getChoice(prompt string,scanner *bufio.Scanner) string {
	var choice string
	for choice != "yes" && choice != "no" {
		fmt.Println(prompt)
		if (scanner.Scan()) {
			choice = strings.TrimSpace(strings.ToLower(scanner.Text()))
		}
		if choice != "yes" && choice != "no" {
			fmt.Println("Choice is invalid, please input choice again")
		}
	}
	return choice
}

func determineChoice(prompt string, scanner *bufio.Scanner) bool {
	choice := getChoice(prompt,scanner)
	return choice == "yes"
}