package forms

type InsertStrategyForm struct {
	Goal         string
	Details      string
	Name         string
	StartedAt    string
	EndedAt      string
	Label        string
	TimeEstimate int
	IsProduction bool
}

type StrategyIdForm struct {
	Id int
}

type StrategyCompletedForm struct {
	Id     int
	Report string
}

type ChangeStrategyStatusForm struct {
	Id     int
	Status string
}
