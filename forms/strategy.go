package forms

import "time"

type StrategyForm struct{}

type InsertStrategyForm struct {
	Goal         string
	Details      string
	Name         string
	StartedAt    time.Time
	EndedAt      time.Time
	Label        string
	Status       string
	TimeEstimate int
	UnitTime     string
	Process      int
	IsProduction bool
}

type StrategyIdForm struct {
	Id int
}

type StrategyCompletedForm struct {
	Id     int
	Report string
}
