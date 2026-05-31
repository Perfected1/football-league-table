package main

import (
	"fmt"
	"football-league-table/league"
)

func main() {
	l := league.NewLeague()

	// Add teams
	l.AddTeam("Arsenal")
	l.AddTeam("Chelsea")
	l.AddTeam("Liverpool")
	l.AddTeam("Manchester City")

	// Matches
	l.RecordMatch("Arsenal", "Chelsea", 2, 1)
	l.RecordMatch("Liverpool", "Manchester City", 1, 1)
	l.RecordMatch("Arsenal", "Liverpool", 3, 0)
	l.RecordMatch("Manchester City", "Chelsea", 2, 0)

	// Display table
	fmt.Println("FOOTBALL LEAGUE TABLE\n")

	fmt.Printf("%-20s %-3s %-3s %-3s %-3s %-4s %-4s %-4s %-4s\n",
		"Team", "P", "W", "D", "L", "GF", "GA", "GD", "Pts")

	for _, t := range l.Standings() {
		fmt.Printf("%-20s %-3d %-3d %-3d %-3d %-4d %-4d %-4d %-4d\n",
			t.Name,
			t.Played,
			t.Won,
			t.Drawn,
			t.Lost,
			t.GoalsFor,
			t.GoalsAgainst,
			t.GoalDifference(),
			t.Points,
		)
	}
}