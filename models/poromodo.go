package models

import (
	"time"

	"github.com/MinhNguyenLe/time-tracking-server/db"
	"github.com/MinhNguyenLe/time-tracking-server/forms"
)

type Poromodo struct {
	Strategy struct {
		Name         string
		Label        string
		Process      int
		TimeEstimate int
		Status       string
		Goal         string
	}
	Satisfaction float64
	Productivity float64
	Interested   float64
	Insight      float64
	CreatedAt    time.Time
	Goal         string
	Duration     int
}

type PoromodoModel struct{}

func (p PoromodoModel) Insert(form forms.InsertPoromodoForm) (poromodoId int64, err error) {
	err = db.GetDB().QueryRow("INSERT INTO poromodo(duration, strategy_id, satisfaction, productivity, interested, insight, created_at, goal) VALUES($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id", form.Duration, form.StrategyId, form.Satisfaction, form.Productivity, form.Interested, form.Insight, time.Now(), form.Goal).Scan(&poromodoId)

	_, errorUpdate := db.GetDB().Query("UPDATE strategy SET process=process+$1, updated_at=$2 WHERE id=$3", form.Duration, time.Now(), form.StrategyId)
	if errorUpdate != nil {
		return 0, errorUpdate
	}

	return poromodoId, err
}

func (s PoromodoModel) GetList() ([]Poromodo, error) {
	rows, err := db.GetDB().Query("SELECT poromodo.goal, duration, poromodo.satisfaction, poromodo.productivity, poromodo.interested, poromodo.insight, strategy.name, strategy.goal, strategy.process, strategy.time_estimate, strategy.label, strategy.status from poromodo inner join strategy on strategy_id=strategy.id order by poromodo.created_at desc	limit 20")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var poromodos []Poromodo

	for rows.Next() {
		var poromodo Poromodo
		if err := rows.Scan(&poromodo.Goal, &poromodo.Duration, &poromodo.Satisfaction, &poromodo.Productivity, &poromodo.Interested, &poromodo.Insight, &poromodo.Strategy.Name, &poromodo.Strategy.Goal, &poromodo.Strategy.Process, &poromodo.Strategy.TimeEstimate, &poromodo.Strategy.Label, &poromodo.Strategy.Status); err != nil {
			return poromodos, err
		}
		poromodos = append(poromodos, poromodo)
	}
	if err = rows.Err(); err != nil {
		return poromodos, err
	}
	return poromodos, nil
}

func (s PoromodoModel) GetByStrategyId(id string) ([]Poromodo, error) {
	rows, err := db.GetDB().Query("SELECT goal, duration, satisfaction, productivity, interested, insight from poromodo WHERE strategy_id=$1 order by created_at desc", id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var poromodos []Poromodo

	for rows.Next() {
		var poromodo Poromodo
		if err := rows.Scan(&poromodo.Goal, &poromodo.Duration, &poromodo.Satisfaction, &poromodo.Productivity, &poromodo.Interested, &poromodo.Insight); err != nil {
			return poromodos, err
		}
		poromodos = append(poromodos, poromodo)
	}
	if err = rows.Err(); err != nil {
		return poromodos, err
	}
	return poromodos, nil
}
