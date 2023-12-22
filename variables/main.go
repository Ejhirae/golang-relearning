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
	fmt.Println("Guess my number, you have 3 guesses")
	fmt.Println("If the difference between guessed number and actual number is in the range of 1 - 25 then you are getting closer, else you are far away")
	fmt.Scan(&userInput)
	guessingGame(guessingNumber, numberOfGuesses, userInput)

}

func guessingGame(guessingNumber int, numberOfGuesses int, userInput int ){

	for userInput != guessingNumber {
		if numberOfGuesses > 1 && math.Abs(float64(userInput-guessingNumber)) <= 25 {
			fmt.Println("Keep guessing you are getting warmer ")
			numberOfGuesses--
			fmt.Println("number of guesses: ", numberOfGuesses)
			fmt.Scan(&userInput)
		} else if numberOfGuesses > 1 && math.Abs(float64(userInput-guessingNumber)) > 25 {
			fmt.Println("Very cold")
			numberOfGuesses--
			fmt.Println("number of guesses: ", numberOfGuesses)
			fmt.Scan(&userInput)
		} else if numberOfGuesses > 1 && userInput == guessingNumber {
			break
		} else {
			fmt.Println("You lost")
			fmt.Println("The correct number was:", guessingNumber)
			os.Exit(0)
		}
	}

	fmt.Println("You won!")

}
