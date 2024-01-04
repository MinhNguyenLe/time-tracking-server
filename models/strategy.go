package models

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/MinhNguyenLe/time-tracking-server/db"
	"github.com/MinhNguyenLe/time-tracking-server/forms"
)

type ListStrategies struct {
	Id           int
	Goal         string
	Details      string
	Name         string
	CreatedAt    time.Time
	StartedAt    time.Time
	EndedAt      time.Time
	Label        string
	Status       string
	TimeEstimate int
	Process      sql.NullInt64
	IsProduction sql.NullBool
}

type StrategyModel struct{}

func (p StrategyModel) Insert(form forms.InsertStrategyForm) (strategyId int64, err error) {
	fmt.Println(form)

	err = db.GetDB().QueryRow(
		"INSERT INTO strategy(name, goal, details, created_at, label, status, time_estimate, started_at, ended_at, unit_time, process, is_production) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) RETURNING id",
		form.Name,
		form.Goal,
		form.Details,
		time.Now(),
		form.Label,
		form.Status,
		form.TimeEstimate,
		form.StartedAt,
		form.EndedAt,
		form.UnitTime,
		form.Process,
		form.IsProduction).Scan(&strategyId)

	return strategyId, err
}

func (s StrategyModel) GetList() ([]ListStrategies, error) {
	rows, err := db.GetDB().Query("SELECT id, name, goal, details, created_at, label, status, time_estimate, started_at, ended_at, process, is_production FROM strategy ORDER BY status desc")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var strategies []ListStrategies

	for rows.Next() {
		var strategy ListStrategies
		if err := rows.Scan(&strategy.Id, &strategy.Name, &strategy.Goal,
			&strategy.Details, &strategy.CreatedAt, &strategy.Label,
			&strategy.Status, &strategy.TimeEstimate,
			&strategy.StartedAt, &strategy.EndedAt,
			&strategy.Process, &strategy.IsProduction); err != nil {
			return strategies, err
		}
		strategies = append(strategies, strategy)
	}
	if err = rows.Err(); err != nil {
		return strategies, err
	}
	return strategies, nil
}

func (s StrategyModel) TriggerInProcess(form forms.StrategyIdForm) error {
	_, err := db.GetDB().Query("SELECT * FROM strategy WHERE id=$1", form.Id)
	if err != nil {
		return err
	}

	_, errorUpdate := db.GetDB().Query("UPDATE strategy SET status=$1 WHERE id=3", "IN_PROCESS", form.Id)
	if errorUpdate != nil {
		return errorUpdate
	}

	return nil
}

func (s StrategyModel) TriggerCompleted(form forms.StrategyIdForm) error {
	_, err := db.GetDB().Query("SELECT * FROM strategy WHERE id=$1", form.Id)
	if err != nil {
		return err
	}

	_, errorUpdate := db.GetDB().Query("UPDATE strategy SET status=$1 WHERE id=$2", "COMPLETED", form.Id)
	if errorUpdate != nil {
		return errorUpdate
	}

	return nil
}
