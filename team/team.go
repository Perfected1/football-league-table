package team

// Team represents a football club in the league.
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