package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"football-league-table/league"
	"football-league-table/team"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter number of teams in the league: ")
	input, _ := reader.ReadString('\n')
	input = team.Normalize(input)

	capacity, err := strconv.Atoi(input)
	if err != nil || capacity <= 0 {
		fmt.Println("Invalid number")
		return
	}

	l := league.NewLeague(capacity)

	// Register teams
	for len(l.Teams) < capacity {
		fmt.Printf("Enter team %d name: ", len(l.Teams)+1)
		name, _ := reader.ReadString('\n')

		err := l.AddTeam(name)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Added:", team.TitleCase(name))
		}
	}

	fmt.Println("\nMATCH INPUT READY (case-insensitive, alias supported)\n")

	for {
		fmt.Print("Enter match: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if strings.ToLower(input) == "exit" {
			break
		}

		parts := strings.Split(input, ",")
		if len(parts) != 4 {
			fmt.Println("Format: Home, Away, HG, AG")
			continue
		}

		home := parts[0]
		away := parts[1]

		hg, e1 := strconv.Atoi(strings.TrimSpace(parts[2]))
		ag, e2 := strconv.Atoi(strings.TrimSpace(parts[3]))

		if e1 != nil || e2 != nil {
			fmt.Println("Goals must be numbers")
			continue
		}

		err := l.RecordMatch(home, away, hg, ag)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		fmt.Println("\nUPDATED TABLE\n")

		fmt.Printf("%-25s %-3s %-3s %-3s %-3s %-4s %-4s %-4s\n",
			"Team", "P", "W", "D", "L", "GD", "GF", "Pts")

		for _, t := range l.Standings() {
	fmt.Printf("%-25s %-3d %-3d %-3d %-3d %-4d %-4d %-4d\n",
		team.TitleCase(t.Name),
		t.Played,
		t.Won,
		t.Drawn,
		t.Lost,
		t.GoalDifference(),
		t.GoalsFor,
		t.Points,
	)
		}

		fmt.Println()
	}
}