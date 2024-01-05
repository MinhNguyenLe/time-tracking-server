package models

import (
	"fmt"
	"time"

	"github.com/MinhNguyenLe/time-tracking-server/db"
	"github.com/MinhNguyenLe/time-tracking-server/forms"
	"github.com/MinhNguyenLe/time-tracking-server/utils"
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
	Process      *int
	IsProduction *bool
}

type ListStrategiesByStatus struct {
	Id      int
	Name    string
	Process *int
	Goal    string
	Label   string
}

type StrategyModel struct{}

func (p StrategyModel) Insert(form forms.InsertStrategyForm) (strategyId int64, err error) {
	fmt.Println(form)

	startedAtParsed, errStartedAt := utils.TimeParsed(form.StartedAt)
	if err != nil {
		return 0, errStartedAt
	}
	endedAtParsed, errEndedAt := utils.TimeParsed(form.EndedAt)
	if err != nil {
		return 0, errEndedAt
	}

	err = db.GetDB().QueryRow(
		"INSERT INTO strategy(name, goal, details, created_at, label, status, time_estimate, started_at, ended_at, process, is_production) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id",
		form.Name,
		form.Goal,
		form.Details,
		time.Now(),
		form.Label,
		form.Status,
		form.TimeEstimate,
		startedAtParsed,
		endedAtParsed,
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

func (s StrategyModel) ChangeStatus(form forms.ChangeStrategyStatusForm) error {
	_, err := db.GetDB().Query("SELECT * FROM strategy WHERE id=$1", form.Id)
	if err != nil {
		return err
	}

	_, errorUpdate := db.GetDB().Query("UPDATE strategy SET status=$1, updated_at=$2 WHERE id=$3", form.Status, time.Now(), form.Id)
	if errorUpdate != nil {
		return errorUpdate
	}

	_, errorSetHistory := db.GetDB().Query("INSERT INTO strategy_changed_history(created_at, status, strategy_id) VALUES($1, $2, $3) RETURNING id", time.Now(), form.Status, form.Id)
	if errorSetHistory != nil {
		return errorSetHistory
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

func (s StrategyModel) GetStrategiesByStatus(form forms.StrategyStatusForm) ([]ListStrategiesByStatus, error) {
	rows, err := db.GetDB().Query("SELECT id, name, process, goal, label FROM strategy WHERE status=$1", form.Status)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var strategies []ListStrategiesByStatus

	for rows.Next() {
		var strategy ListStrategiesByStatus
		if err := rows.Scan(&strategy.Id, &strategy.Name, &strategy.Process, &strategy.Goal, &strategy.Label); err != nil {
			return strategies, err
		}
		strategies = append(strategies, strategy)
	}
	if err = rows.Err(); err != nil {
		return strategies, err
	}
	return strategies, nil
}
