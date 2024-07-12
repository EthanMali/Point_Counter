package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Player struct {
	Name  string
	Score int
}

var players []Player

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("1. Add Player")
		fmt.Println("2. Update Score")
		fmt.Println("3. Display Scores")
		fmt.Println("4. Exit")
		fmt.Print("Enter your choice: ")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			addPlayer(scanner)
		case "2":
			updateScore(scanner)
		case "3":
			displayScores()
		case "4":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func addPlayer(scanner *bufio.Scanner) {
	fmt.Print("Enter player name: ")
	scanner.Scan()
	name := scanner.Text()
	players = append(players, Player{Name: name, Score: 0})
	fmt.Println("Player added successfully.")
}

func updateScore(scanner *bufio.Scanner) {
	fmt.Print("Enter player name: ")
	scanner.Scan()
	name := scanner.Text()

	for i, player := range players {
		if player.Name == name {
			fmt.Print("Enter score to add: ")
			scanner.Scan()
			scoreInput := scanner.Text()
			score, err := strconv.Atoi(scoreInput)
			if err != nil {
				fmt.Println("Invalid score. Please enter a number.")
				return
			}
			players[i].Score += score
			fmt.Println("Score updated successfully.")
			return
		}
	}

	fmt.Println("Player not found.")
}

func displayScores() {
	fmt.Println("Current Scores:")
	for _, player := range players {
		fmt.Printf("%s: %d\n", player.Name, player.Score)
	}
}
