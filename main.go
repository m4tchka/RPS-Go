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
	//NB empty strings not accepted - will only move on once at least 1 character is entered
	return playerName
	//Return the playerName variable
}
func getPlayerMove() string {
	var playerInput, playerMove string
	//Declare 2 string variables
	fmt.Print("Input a move: ")
	//Print the string above to console
	fmt.Scan(&playerInput)
	//Takes the user's input and (effectively) assigns it to the variable playerInput
	playerInput = strings.ToLower(playerInput)
	//Turn the player's input into lowercase
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
	// Assign the player's input to the variable playerMove
	// If the input is not a valid move,  print the string "Invalid move!" and reprompt
	return playerMove
}
func getComputerMove() string {
	var computerMove string
	rand.Seed(time.Now().UnixNano())
	// Changes RNG seed to be current time at function call from Unix time, in nanoseconds
	number := rand.Intn(3)
	// Rolls an INTEGER between 0 (inclusive) and 3 (exclusive) - A half-open (or half-closed) interval of numbers, notation: [0, n)
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
	// Assign a move to computerMove according to the number rolled, and print it to the terminal. Then return said move
	return computerMove
}
func findResult(move, cMove string) string {
	var result string
	if move == "stop" {
		result = "Stopping..."
		// First check if the move was "stop" - assign the result variable the string "Stopping..."" (which is printed in the console by line 101 when findResult is called)"
	} else if move == cMove {
		result = "Draw"
		playerD++
		// Next, check if the player's move matches the computer's random move - if it is, increment the draw count
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
		// 2 x else if blocks, checking who won and incrementing the score accordingly; they may only be triggered if the move was "rock" or "paper" or "scissors"
	} else {
		result = "Please input a valid move: rock, paper or scissors, or type 'stop' to exit the game !"
	}
	// To get to this else statement, all the other checks must have been false - therefore the move is invalid !
	// Assign to the result variable a string telling the user the valid inputs (which is printed in the console by line 101 when findResult is called), and does not increment the score.
	return result
}

var playerW, playerL, playerD int

// Declare score-keeping variables

func main() {
	var name, move, cMove string
	name = getPlayerName()
	// Assign the returned variable from getPlayerName to the variable name
	fmt.Printf("Greetings, %v. \n", name)
	// Greet the player
	for move != "stop" {
		// Whilst the move variable is not the string "stop", do the following:
		move, cMove = getPlayerMove(), getComputerMove()
		// Assign the move and cMove variables to be the values returned from the 2 functions above
		fmt.Println(findResult(move, cMove))
		// Print the result of handing the previous variables to the findResult function
		fmt.Printf("Current score for %v is: Wins: %d -- Draws %d -- Losses: %d \n", name, playerW, playerD, playerL)
		// Print the current score for the player
	}
	if move == "stop" {
		// If the move variable does have the value "stop" then:
		fmt.Printf("Have a good day, %v!", name)
		// Say a farewell to the player, and conclude the function main function, exiting the program !
	}
}
