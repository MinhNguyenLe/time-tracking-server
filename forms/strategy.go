package forms

import "time"

type StrategyForm struct{}

type InsertStrategyForm struct {
	Goal         string
	Details      string
	Name         string
	CreatedAt    time.Time
	StartedAt    time.Time
	EndedAt      time.Time
	Label        string
	Status       string
	TimeEstimate int
	UnitTime     string
	Process      int
	IsProduction bool
}
