package league

import (
	"errors"
	"football-league-table/team"
	"sort"
	"strings"
)

// League uses map-based storage (stable references)
type League struct {
	Teams    map[string]*team.Team
	Order    []string
	Capacity int
}

// NewLeague creates league
func NewLeague(capacity int) League {
	return League{
		Teams:    make(map[string]*team.Team),
		Order:    []string{},
		Capacity: capacity,
	}
}

// normalize input
func normalize(s string) string {
	return strings.ToUpper(strings.ReplaceAll(strings.TrimSpace(s), " ", ""))
}

// AddTeam registers team
func (l *League) AddTeam(name string) error {
	if len(l.Teams) >= l.Capacity {
		return errors.New("league is full")
	}

	key := normalize(name)

	if _, exists := l.Teams[key]; exists {
		return errors.New("team already exists: " + name)
	}

	t := &team.Team{
		Name:  name,
		Alias: team.GenerateAlias(name),
	}

	l.Teams[key] = t
	l.Order = append(l.Order, key)

	return nil
}

// resolve team safely
func (l *League) resolve(input string) (*team.Team, error) {
	key := normalize(input)

	if t, ok := l.Teams[key]; ok {
		return t, nil
	}

	// alias fallback
	for _, t := range l.Teams {
		if normalize(t.Alias) == key {
			return t, nil
		}
	}

	return nil, errors.New("unknown team: " + input)
}

// RecordMatch updates stats
func (l *League) RecordMatch(home, away string, hg, ag int) error {
	if normalize(home) == normalize(away) {
		return errors.New("team cannot play itself")
	}

	homeTeam, err := l.resolve(home)
	if err != nil {
		return err
	}

	awayTeam, err := l.resolve(away)
	if err != nil {
		return err
	}

	homeTeam.Played++
	awayTeam.Played++

	homeTeam.GoalsFor += hg
	homeTeam.GoalsAgainst += ag

	awayTeam.GoalsFor += ag
	awayTeam.GoalsAgainst += hg

	if hg > ag {
		homeTeam.Won++
		homeTeam.Points += 3
		awayTeam.Lost++
	} else if ag > hg {
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

// Standings returns sorted table
func (l *League) Standings() []*team.Team {
	list := []*team.Team{}

	for _, t := range l.Teams {
		list = append(list, t)
	}

	sort.Slice(list, func(i, j int) bool {
		if list[i].Points == list[j].Points {
			if list[i].GoalDifference() == list[j].GoalDifference() {
				return list[i].GoalsFor > list[j].GoalsFor
			}
			return list[i].GoalDifference() > list[j].GoalDifference()
		}
		return list[i].Points > list[j].Points
	})

	return list
}