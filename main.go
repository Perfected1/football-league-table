package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"football-league-table/league"
)

func main() {
	l := league.NewLeague()

	// Add teams (now with validation)
	teams := []string{"Arsenal", "Chelsea", "Liverpool", "Manchester City"}

	for _, t := range teams {
		err := l.AddTeam(t)
		if err != nil {
			fmt.Println("Error adding team:", err)
		}
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("FOOTBALL LEAGUE CLI")
	fmt.Println("Enter match results in format:")
	fmt.Println("HomeTeam, AwayTeam, HomeGoals, AwayGoals")
	fmt.Println("Type 'exit' to stop\n")

	for {
		fmt.Print("Enter match: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "exit" {
			break
		}

		parts := strings.Split(input, ",")
		if len(parts) != 4 {
			fmt.Println("Invalid format. Use: Home, Away, HG, AG")
			continue
		}

		home := strings.TrimSpace(parts[0])
		away := strings.TrimSpace(parts[1])

		homeGoals, err1 := strconv.Atoi(strings.TrimSpace(parts[2]))
		awayGoals, err2 := strconv.Atoi(strings.TrimSpace(parts[3]))

		if err1 != nil || err2 != nil {
			fmt.Println("Goals must be valid numbers.")
			continue
		}

		err := l.RecordMatch(home, away, homeGoals, awayGoals)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		// Print updated table
		fmt.Println("\nUPDATED TABLE\n")

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

		fmt.Println()
	}
}