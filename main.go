package main

import "fmt"

// Team stores information about a football club.
type Team struct {
	Name         string
	Played       int
	Won          int
	Drawn        int
	Lost         int
	GoalsFor     int
	GoalsAgainst int
	Points       int
}

func main() {
	// Create a sample team.
	arsenal := Team{
		Name: "Arsenal",
	}

	fmt.Println("Football League Table")
	fmt.Println("Team:", arsenal.Name)
	fmt.Println("Points:", arsenal.Points)
}