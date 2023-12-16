package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
)

// building a guessing game
// getting warmer or colder
// Checks if the difference between guessed number and actual number is 25
func main() {
	guessingNumber := rand.Intn(98) + 2
	numberOfGuesses := 3
	var userInput int
	fmt.Println("Guess a Number between 1 - 100")
	fmt.Scan(&userInput)

	for userInput != guessingNumber {
		if numberOfGuesses > 1 && math.Abs(float64(userInput-guessingNumber)) <= 25 {
			fmt.Println("Keep guessing you are getting warmer ")
			numberOfGuesses--
			fmt.Println("number of guesses: ", numberOfGuesses)
			fmt.Scan(&userInput)
		} else if numberOfGuesses > 0 && math.Abs(float64(userInput-guessingNumber)) > 25 {
			fmt.Println("Very cold")
			numberOfGuesses--
			fmt.Println("number of guesses: ", numberOfGuesses)
			fmt.Scan(&userInput)
		} else if numberOfGuesses > 0 && userInput == guessingNumber {
			break
		} else {
			fmt.Println("You lost")
			fmt.Println("The correct number was:", guessingNumber)
			os.Exit(0)
		}
	}

	fmt.Println("You won!")
}
