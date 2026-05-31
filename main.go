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

// findTeam returns a pointer to a team in the slice.
func findTeam(teams []Team, name string) *Team {
	for i := range teams {
		if teams[i].Name == name {
			return &teams[i]
		}
	}
	return nil
}

// recordMatch updates team stats based on match result.
func recordMatch(teams []Team, home string, away string, homeGoals, awayGoals int) {
	homeTeam := findTeam(teams, home)
	awayTeam := findTeam(teams, away)

	if homeTeam == nil || awayTeam == nil {
		return
	}

	// Update played games
	homeTeam.Played++
	awayTeam.Played++

	// Update goals
	homeTeam.GoalsFor += homeGoals
	homeTeam.GoalsAgainst += awayGoals

	awayTeam.GoalsFor += awayGoals
	awayTeam.GoalsAgainst += homeGoals

	// Decide result
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

	// Record a match result
	recordMatch(teams, "Arsenal", "Chelsea", 2, 1)

	fmt.Println("FOOTBALL LEAGUE TABLE")
	fmt.Println()

	fmt.Printf("%-20s %-3s %-3s %-3s %-3s %-4s %-4s %-4s\n",
		"Team", "P", "W", "D", "L", "GF", "GA", "Pts")

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