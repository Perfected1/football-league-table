package league

import (
	"errors"
	"football-league-table/team"
	"sort"
)

// League holds all teams with a fixed capacity.
type League struct {
	Teams    []team.Team
	Capacity int
}

// NewLeague creates a league with a fixed number of teams allowed.
func NewLeague(capacity int) League {
	return League{
		Teams:    []team.Team{},
		Capacity: capacity,
	}
}

// AddTeam adds a team if capacity is not exceeded and no duplicates exist.
func (l *League) AddTeam(name string) error {
	// Check duplicate
	for _, t := range l.Teams {
		if t.Name == name {
			return errors.New("team already exists: " + name)
		}
	}

	// Check capacity limit
	if len(l.Teams) >= l.Capacity {
		return errors.New("league is full. maximum teams allowed reached")
	}

	l.Teams = append(l.Teams, team.Team{Name: name})
	return nil
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
func (l *League) RecordMatch(home, away string, homeGoals, awayGoals int) error {
	homeTeam := l.findTeam(home)
	awayTeam := l.findTeam(away)

	if home == away {
		return errors.New("a team cannot play itself")
	}

	if homeTeam == nil {
		return errors.New("unknown team: " + home)
	}
	if awayTeam == nil {
		return errors.New("unknown team: " + away)
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

	return nil
}

// Standings returns sorted table.
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