// guessNumber is a game where the user tries to guess a randomly generated
// number
package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	target := rand.Int63n(100) + 1
	reader := bufio.NewReader(os.Stdin)
	guessesLeft := 10
	fmt.Print("Try to guess my number between 1 - 100\n")
	for guessesLeft > 0 {
		guess := getGuess(reader)
		gameStatus := target - guess
		if gameStatus == 0 {
			fmt.Println("You got it!")
			return
		} else {
			problem := "Too Low!"
			if gameStatus < 0 {
				problem = "Too High!"
			}
			fmt.Println(problem)
			guessesLeft--
		}
	}
	fmt.Println("No guesses left!\nThe answer was", target)
}

func getGuess(reader *bufio.Reader) int64 {
	fmt.Print("Enter your next guess:\n> ")
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	guess, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	return guess
}
