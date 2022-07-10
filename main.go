package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func getPlayerName() string {
	var playerName string
	fmt.Scan(&playerName)
	return playerName
}
func getComputerMove() string {
	var computerMove string
	rand.Seed(time.Now().UnixNano())
	number := rand.Intn(3)
	switch number {
	case 0:
		computerMove = "rock"
		fmt.Printf("The computer chose %v.", computerMove)
	case 1:
		computerMove = "paper"
		fmt.Printf("The computer chose %v.", computerMove)
	default:
		computerMove = "scissors"
		fmt.Printf("The computer chose %v.", computerMove)
	}
	return computerMove
}
func getPlayerMove() string {
	var playerInput, playerMove string
	fmt.Println("Please input a move")
	fmt.Scan(&playerInput)
	if playerInput != "" {
		strings.ToLower(playerInput)
		switch playerInput {
		case "rock":
			playerMove = "rock"
		case "paper":
			playerMove = "paper"
		case "scissors":
			playerMove = "scissors"
		default:
			fmt.Println("Invalid move!")
		}
	} else {
		fmt.Println("Please input a move!")
	}
	return playerMove
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
	fmt.Println("Please enter your name:")
	var name, move, cMove string
	name = getPlayerName()
	fmt.Printf("Greetings, %v.", name)
	move = getPlayerMove()
	cMove = getComputerMove()
	fmt.Println(findResult(move, cMove))
	fmt.Printf("Current score for %v is:", name)
	fmt.Printf("Wins: %d -- Draws %d -- Losses: %d", playerW, playerD, playerL)
}
