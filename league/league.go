package league

import (
	"football-league-table/team"
	"sort"
)

// League holds all teams.
type League struct {
	Teams []team.Team
}

// NewLeague creates a new league.
func NewLeague() League {
	return League{
		Teams: []team.Team{},
	}
}

// AddTeam adds a team to the league.
func (l *League) AddTeam(name string) {
	l.Teams = append(l.Teams, team.Team{Name: name})
}

// findTeam returns pointer to a team.
func (l *League) findTeam(name string) *team.Team {
	for i := range l.Teams {
		if l.Teams[i].Name == name {
			return &l.Teams[i]
		}
	}
	return nil
}

// RecordMatch updates stats after a match.
func (l *League) RecordMatch(home, away string, homeGoals, awayGoals int) {
	homeTeam := l.findTeam(home)
	awayTeam := l.findTeam(away)

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

// Standings returns sorted league table.
func (l *League) Standings() []team.Team {
	sorted := make([]team.Team, len(l.Teams))
	copy(sorted, l.Teams)

	sort.Slice(sorted, func(i, j int) bool {
		if sorted[i].Points == sorted[j].Points {
			if sorted[i].GoalDifference() == sorted[j].GoalDifference() {
				return sorted[i].GoalsFor > sorted[j].GoalsFor
			}
			return sorted[i].GoalDifference() > sorted[j].GoalDifference()
		}
		return sorted[i].Points > sorted[j].Points
	})

	return sorted
}