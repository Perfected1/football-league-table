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
	reader := bufio.NewReader(os.Stdin)

	// Ask for league size
	fmt.Print("Enter number of teams in the league: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	capacity, err := strconv.Atoi(input)
	if err != nil || capacity <= 0 {
		fmt.Println("Invalid number of teams")
		return
	}

	l := league.NewLeague(capacity)

	fmt.Printf("\nYou can now enter %d teams\n\n", capacity)

	// Register teams
	for len(l.Teams) < capacity {
		fmt.Printf("Enter team %d name: ", len(l.Teams)+1)
		name, _ := reader.ReadString('\n')
		name = strings.TrimSpace(name)

		err := l.AddTeam(name)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Added:", name)
		}
	}

	fmt.Println("\nLEAGUE FULL. You can now enter matches.\n")

	fmt.Println("Match format:")
	fmt.Println("HomeTeam, AwayTeam, HomeGoals, AwayGoals")
	fmt.Println("Type 'exit' to stop\n")

	// Match input loop
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
			fmt.Println("Goals must be numbers")
			continue
		}

		err := l.RecordMatch(home, away, homeGoals, awayGoals)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		// Print table
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