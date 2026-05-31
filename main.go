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
	// Create a list of teams in the league.
	teams := []Team{
		{Name: "Arsenal"},
		{Name: "Chelsea"},
		{Name: "Liverpool"},
		{Name: "Manchester City"},
	}

	fmt.Println("Football League Table")
	fmt.Println()

	// Show all teams currently in the league.
	for _, team := range teams {
		fmt.Println(team.Name)
	}
}