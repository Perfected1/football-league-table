package main

import (
	"fmt"
	"sort"
)

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

// GoalDifference calculates GD dynamically.
func (t Team) GoalDifference() int {
	return t.GoalsFor - t.GoalsAgainst
}

// findTeam returns a pointer to a team in the slice.
func findTeam(teams []Team, name string) *Team {
	for i := range teams {
		if teams[i].Name == name {
			return &teams[i]
		}
	}
	return nil
}

// recordMatch updates stats after a match.
func recordMatch(teams []Team, home, away string, homeGoals, awayGoals int) {
	homeTeam := findTeam(teams, home)
	awayTeam := findTeam(teams, away)

	if homeTeam == nil || awayTeam == nil {
		return
	}

	homeTeam.Played++
	awayTeam.Played++

	homeTeam.GoalsFor += homeGoals
	homeTeam.GoalsAgainst += awayGoals

	awayTeam.GoalsFor += awayGoals
	awayTeam.GoalsAgainst += homeGoals

	if homeGoals > awayGoals {
		homeTeam.Won++
		homeTeam.Points += 3
		awayTeam.Lost++
	} else if homeGoals < awayGoals {
		awayTeam.Won++
		awayTeam.Points += 3
		homeTeam.Lost++
	} else {
		homeTeam.Drawn++
		awayTeam.Drawn++
		homeTeam.Points++
		awayTeam.Points++
	}
}

func main() {
	teams := []Team{
		{Name: "Arsenal"},
		{Name: "Chelsea"},
		{Name: "Liverpool"},
		{Name: "Manchester City"},
	}

	// Sample fixtures to populate the table.
	recordMatch(teams, "Arsenal", "Chelsea", 2, 1)
	recordMatch(teams, "Liverpool", "Manchester City", 1, 1)
	recordMatch(teams, "Arsenal", "Liverpool", 3, 0)
	recordMatch(teams, "Manchester City", "Chelsea", 2, 0)

	// Sort league table
	sort.Slice(teams, func(i, j int) bool {
		if teams[i].Points == teams[j].Points {
			if teams[i].GoalDifference() == teams[j].GoalDifference() {
				return teams[i].GoalsFor > teams[j].GoalsFor
			}
			return teams[i].GoalDifference() > teams[j].GoalDifference()
		}
		return teams[i].Points > teams[j].Points
	})

	fmt.Println("FOOTBALL LEAGUE TABLE")
	fmt.Println()

	fmt.Printf("%-20s %-3s %-3s %-3s %-3s %-4s %-4s %-4s %-4s\n",
		"Team", "P", "W", "D", "L", "GF", "GA", "GD", "Pts")

	for _, team := range teams {
		fmt.Printf("%-20s %-3d %-3d %-3d %-3d %-4d %-4d %-4d %-4d\n",
			team.Name,
			team.Played,
			team.Won,
			team.Drawn,
			team.Lost,
			team.GoalsFor,
			team.GoalsAgainst,
			team.GoalDifference(),
			team.Points,

		)
	}
}