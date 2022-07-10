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
func getPlayerMove() string {
	var playerInput, playerMove string
	fmt.Println("Please input a move")
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
		default:
			fmt.Println("Invalid move!")
		}
	} else {
		fmt.Println("Please input a move!")
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
		fmt.Println(fmt.Sprintf("The computer chose %v.", computerMove))
	case 1:
		computerMove = "paper"
		fmt.Println(fmt.Sprintf("The computer chose %v.", computerMove))
	default:
		computerMove = "scissors"
		fmt.Println(fmt.Sprintf("The computer chose %v.", computerMove))
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
	fmt.Println("Please enter your name:")
	name = getPlayerName()
	fmt.Println(fmt.Sprintf("Greetings, %v.", name))
	move, cMove = getPlayerMove(), getComputerMove()
	fmt.Println(findResult(move, cMove))
	fmt.Println(fmt.Sprintf("Current score for %v is:", name))
	fmt.Println(fmt.Sprintf("Wins: %d -- Draws %d -- Losses: %d", playerW, playerD, playerL))
}
