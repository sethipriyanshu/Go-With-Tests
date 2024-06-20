package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strings"
)

func main() {
	fmt.Println("Welcome to Rock Paper Scissors! Type quit to exit")
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Type R for rock, P for paper and S for scissors: ")
		userChoice, err := reader.ReadString('\n')
		userChoice = strings.TrimSpace(strings.ToLower(userChoice))
		if err != nil {
			fmt.Println("An Unexpected Error Has Occured")
			break
		}
		if userChoice == "quit" {
			break
		}
		compChoice := computerChoice()
		if userChoice == "R" || userChoice == "r" {
			if compChoice == "rock" {
				fmt.Println("You Both Chose Rock. Try Again!")
			}
			if compChoice == "paper" {
				fmt.Println("Opponent Chose Paper, You Lose!")
			}
			if compChoice == "scissors" {
				fmt.Println("Opponent Chose Scissors, You Win!")

			}
		} else if userChoice == "P" || userChoice == "p" {
			if compChoice == "rock" {
				fmt.Println("Opponent Chose Rock, You Win!")
			}
			if compChoice == "paper" {
				fmt.Println("You both chose paper, Try Again!")
			}
			if compChoice == "scissors" {
				fmt.Println("Opponent Chose Scissors, You Lose!")
			}
		} else if userChoice == "S" || userChoice == "s" {
			if compChoice == "rock" {
				fmt.Println("Opponent Chose Rock, You Lose!")
			}
			if compChoice == "paper" {
				fmt.Println("Opponent Chose Paper, You Win!")
			}
			if compChoice == "scissors" {
				fmt.Println("You both chose Scissors, Try Again!")
			}
		} else {
			fmt.Println("Invalid Input, Try Again!")
		}

	}

}

func computerChoice() string {
	intChoice := rand.IntN(3)
	var choice string
	switch intChoice {
	case 0:
		choice = "rock"
	case 1:
		choice = "paper"
	case 2:
		choice = "scissors"
	}
	return choice
}
