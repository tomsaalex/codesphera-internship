package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
)

/*
Guess the Number Game

Problem:
The program randomly selects a number between 1 and 100.
The user must guess the number in a limited number of tries (variable), receiving feedback:
- "Too low!" if the guess is below the number
- "Too high!" if the guess is above the number
- "Correct!" if the guess is right

If the user has lost, return "You have lost!" every time a guess is tried
If the user won, return "You have won!" every time a guess is tried

Input: integers from user input (guesses)
Output: feedback strings

Implement the Guess function that would pass the tests, as well as a
command line user interface that would generate a game (with user given
max retries), generate a secret random number, then let the user play the game.
*/

func generateSecretNumber() int {
	return rand.IntN(100) + 1
}

func readPositiveNumber() int {
	inputSuccessful := false
	var num int

	reader := bufio.NewReader(os.Stdin)

	for !inputSuccessful {
		inputSuccessful = true
		_, err := fmt.Scanf("%d", &num)

		if err != nil {
			inputSuccessful = false
			reader.ReadString('\n')
			fmt.Printf("Whoops! Your input doesn't seem to be valid. Take a look at what's wrong and try again\nError: %s\n", err.Error())
		}

		if num <= 0 {
			inputSuccessful = false
			fmt.Printf("Whoops! Looks like you entered a negative number, but we are expecting a number greater than 0. Try again!\n")
		}
	}

	return num
}

func run_game() {
	fmt.Println("Hello! Welcome to Guess the Number! I will generate a random number betwwen 1 and 100 and you will have to guess, based on the clues provided!")
	fmt.Println("Before we begin, please enter the number of attempts you would like to have at guessing the number:")

	maxAttempts := readPositiveNumber()

	fmt.Printf("Great! You will be given %d attempts to guess the number!\n", maxAttempts)

	game := NewGame(generateSecretNumber(), maxAttempts)

	for {
		fmt.Printf("Please enter your guess: (%d/%d)\n", game.Tries+1, game.MaxTries)
		guessedNumber := readPositiveNumber()
		fmt.Println(game.Guess(guessedNumber))
	}
}

func main() {
	run_game()
}
