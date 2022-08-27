package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func getPlayerName() string {
	var playerName string
	//Declare a variable called playerName, that takes a string as value, but don't assign a value yet.
	fmt.Println("Please enter your name: ")
	//Print the string above to console
	fmt.Scan(&playerName)
	//Takes whatever the user types into the console and presses enter on, and assigns it as the value associated with the address of the variable, playerName.
	return playerName
	//Return the variable
}
func getPlayerMove() string {
	var playerInput, playerMove string
	//Declare 2 string variables
	fmt.Print("Please input a move, or type stop to end the game: ")
	//Print the string above to console
	fmt.Scan(&playerInput)
	//Takes the user's input and (effectively) assigns it to variable playerInput
	playerInput = strings.ToLower(playerInput)
	//Lowercase the variable playerInput
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
	if move == "stop" {
		result = "Stopping..."
	} else if move == cMove {
		result = "Draw"
		playerD++
	} else if (move == "rock" && cMove == "scissors") ||
		(move == "paper" && cMove == "rock") ||
		(move == "scissors" && cMove == "paper") {
		result = "Player wins!"
		playerW++
	} else if (move == "rock" && cMove == "paper") ||
		(move == "paper" && cMove == "scissors") ||
		(move == "scissors" && cMove == "rock") {
		result = "Computer wins!"
		playerL++
	} else {
		fmt.Println("Please input a valid move: rock, paper or scissors,")
		fmt.Println("or type 'stop' to exit the game !")
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

// random string input
// if move = stop, stop
// if move = cmove, draw
// if move = r/p/s & cmove = s/r/p, win/lose (and v,v)
// else reprompt
