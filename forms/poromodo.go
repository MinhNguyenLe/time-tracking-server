package forms

type PoromodoForm struct{}

type InsertPoromodoForm struct {
	Duration     int
	StrategyId   int
	Satisfaction int
	Productivity int
	Interested   int
	Insight      int
	Goal         string
}
