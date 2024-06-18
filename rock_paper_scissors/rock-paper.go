package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	fmt.Println("Welcome to Rock Paper Scissors! Type quit to exit")

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
