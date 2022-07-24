package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func getPlayerName() string {
	var playerName string
	fmt.Print("Please enter your name: ")
	fmt.Scan(&playerName)
	return playerName
}
func getPlayerMove() string {
	var playerInput, playerMove string
	fmt.Print("Please input a move: ")
	fmt.Scan(&playerInput)
	playerInput = strings.ToLower(playerInput)
	if playerInput != "" {
		switch playerInput {
		case "rock":
			playerMove = "rock"
		case "paper":
			playerMove = "paper"
		case "scissors":
			playerMove = "scissors"
		case "stop":
			playerMove = "stop"
		default:
			fmt.Println("Invalid move!")
		}
	}
	return playerMove
}
func getComputerMove() string {
	var computerMove string
	rand.Seed(time.Now().UnixNano())
	number := rand.Intn(3)
	switch number {
	case 0:
		computerMove = "rock"
		fmt.Printf("The computer chose %v. \n", computerMove)
	case 1:
		computerMove = "paper"
		fmt.Printf("The computer chose %v. \n", computerMove)
	case 2:
		computerMove = "scissors"
		fmt.Printf("The computer chose %v. \n", computerMove)
	}
	return computerMove
}
func findResult(move, cMove string) string {
	var result string
	if move != cMove {
		if (move == "rock" && cMove == "scissors") ||
			(move == "paper" && cMove == "rock") ||
			(move == "scissors" && cMove == "paper") {
			result = "Player wins!"
			playerW++
		} else {
			result = "Computer wins!"
			playerL++
		}
	} else {
		result = "Draw"
		playerD++
	}
	return result
}

var playerW, playerL, playerD int

func main() {
	var name, move, cMove string
	name = getPlayerName()
	fmt.Printf("Greetings, %v. \n", name)
	for move != "stop" {
		move, cMove = getPlayerMove(), getComputerMove()
		fmt.Println(findResult(move, cMove))
		fmt.Printf("Current score for %v is: Wins: %d -- Draws %d -- Losses: %d \n", name, playerW, playerD, playerL)
	}
	if move == "stop" {
		fmt.Printf("Have a good day, %v!", name)
	}

}
