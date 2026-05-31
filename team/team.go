package team

import (
	"strings"
)

// Team represents a football club.
type Team struct {
	Name         string
	Alias        string
	Played       int
	Won          int
	Drawn        int
	Lost         int
	GoalsFor     int
	GoalsAgainst int
	Points       int
}

// GoalDifference returns GD.
func (t Team) GoalDifference() int {
	return t.GoalsFor - t.GoalsAgainst
}

// GenerateAlias creates shorthand for input matching.
func GenerateAlias(name string) string {
	words := strings.Fields(strings.ToUpper(name))

	if len(words) == 1 {
		if len(words[0]) >= 3 {
			return words[0][:3]
		}
		return words[0]
	}

	first := words[0]
	second := words[1]

	alias := ""
	if len(first) >= 3 {
		alias = first[:3]
	} else {
		alias = first
	}

	alias += " " + string(second[0])
	return alias
}

// Normalize input (case-insensitive matching)
func Normalize(input string) string {
	return strings.ToUpper(strings.TrimSpace(input))
}

// TitleCase formats team names for display
func TitleCase(name string) string {
	words := strings.Fields(strings.ToLower(strings.TrimSpace(name)))

	for i, w := range words {
		if len(w) > 0 {
			words[i] = strings.ToUpper(string(w[0])) + w[1:]
		}
	}

	return strings.Join(words, " ")
}