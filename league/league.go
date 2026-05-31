package league

import (
	"errors"
	"football-league-table/team"
	"sort"
	"strings"
)

// League engine
type League struct {
	Teams    map[string]*team.Team
	Capacity int

	fixtures   map[string]int
	maxGames   int
	finished   bool
	winner     *team.Team
}

// NewLeague creates league with computed limits
func NewLeague(capacity int) League {
	return League{
		Teams:    make(map[string]*team.Team),
		Capacity: capacity,
		fixtures: make(map[string]int),
		maxGames: (capacity - 1) * 2,
	}
}

// canonical ID
func canon(s string) string {
	return strings.ToUpper(strings.TrimSpace(s))
}

// alias normalization
func aliasKey(s string) string {
	return strings.ToUpper(strings.ReplaceAll(strings.TrimSpace(s), " ", ""))
}

// fixture key (order independent)
func fixtureKey(a, b string) string {
	if a < b {
		return a + "|" + b
	}
	return b + "|" + a
}

// AddTeam
func (l *League) AddTeam(name string) error {
	if len(l.Teams) >= l.Capacity {
		return errors.New("league is full")
	}

	id := canon(name)

	if _, exists := l.Teams[id]; exists {
		return errors.New("team already exists: " + name)
	}

	l.Teams[id] = &team.Team{
		Name:  name,
		Alias: team.GenerateAlias(name),
	}

	return nil
}

// resolve team
func (l *League) resolve(input string) (*team.Team, string, error) {
	raw := strings.TrimSpace(input)

	if t, ok := l.Teams[canon(raw)]; ok {
		return t, canon(raw), nil
	}

	for id, t := range l.Teams {
		if aliasKey(t.Alias) == aliasKey(raw) {
			return t, id, nil
		}
	}

	for id, t := range l.Teams {
		if strings.EqualFold(t.Name, raw) {
			return t, id, nil
		}
	}

	return nil, "", errors.New("unknown team: " + input)
}

// check if league ended
func (l *League) checkFinished() {
	for _, t := range l.Teams {
		if t.Played < l.maxGames {
			return
		}
	}

	l.finished = true

	// determine winner
	var top *team.Team

	for _, t := range l.Teams {
		if top == nil ||
			t.Points > top.Points ||
			(t.Points == top.Points && t.GoalDifference() > top.GoalDifference()) {
			top = t
		}
	}

	l.winner = top
}

// RecordMatch with full constraints
func (l *League) RecordMatch(home, away string, hg, ag int) error {
	if l.finished {
		return errors.New("league has ended")
	}

	if strings.EqualFold(home, away) {
		return errors.New("team cannot play itself")
	}

	homeTeam, homeID, err := l.resolve(home)
	if err != nil {
		return err
	}

	awayTeam, awayID, err := l.resolve(away)
	if err != nil {
		return err
	}

	// per-team limit check
	if homeTeam.Played >= l.maxGames || awayTeam.Played >= l.maxGames {
		return errors.New("one of the teams has reached maximum games")
	}

	fKey := fixtureKey(homeID, awayID)

	if l.fixtures[fKey] >= 2 {
		return errors.New("this fixture already played twice")
	}

	l.fixtures[fKey]++

	// update stats
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

	l.checkFinished()

	return nil
}

// Standings
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

// IsFinished returns league status
func (l *League) IsFinished() bool {
	return l.finished
}

// Winner returns champion
func (l *League) Winner() *team.Team {
	return l.winner
}

// MaxGames returns limit per team
func (l *League) MaxGames() int {
	return l.maxGames
}