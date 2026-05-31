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
	// Teams participating in the league.
	teams := []Team{
		{Name: "Arsenal"},
		{Name: "Chelsea"},
		{Name: "Liverpool"},
		{Name: "Manchester City"},
	}

	fmt.Println("FOOTBALL LEAGUE TABLE")
	fmt.Println()

	// Print table header.
	fmt.Printf("%-20s %-3s %-3s %-3s %-3s %-4s %-4s %-4s\n",
		"Team", "P", "W", "D", "L", "GF", "GA", "Pts")

	// Print a line for each team.
	for _, team := range teams {
		fmt.Printf("%-20s %-3d %-3d %-3d %-3d %-4d %-4d %-4d\n",
			team.Name,
			team.Played,
			team.Won,
			team.Drawn,
			team.Lost,
			team.GoalsFor,
			team.GoalsAgainst,
			team.Points,
		)
	}
}