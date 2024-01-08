package models

import (
	"fmt"
	"time"

	"github.com/MinhNguyenLe/time-tracking-server/db"
	"github.com/MinhNguyenLe/time-tracking-server/forms"
	"github.com/MinhNguyenLe/time-tracking-server/utils"
)

type Strategy struct {
	Id           int
	Goal         string
	Details      string
	Name         string
	CreatedAt    time.Time
	StartedAt    time.Time
	EndedAt      time.Time
	Label        string
	Status       string
	TimeEstimate float64
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

type ChangedHistory struct {
	CreatedAt time.Time
	Status    string
}

type StrategyModel struct{}

func (p StrategyModel) Insert(form forms.InsertStrategyForm) (strategyId int64, err error) {
	startedAtParsed, errStartedAt := utils.TimeParsed(form.StartedAt)
	if err != nil {
		return 0, errStartedAt
	}
	endedAtParsed, errEndedAt := utils.TimeParsed(form.EndedAt)
	if err != nil {
		return 0, errEndedAt
	}
	fmt.Print(form.StartedAt)

	fmt.Print(startedAtParsed)
	fmt.Print(endedAtParsed)

	err = db.GetDB().QueryRow(
		"INSERT INTO strategy(name, goal, details, created_at, label, status, time_estimate, started_at, ended_at, process, is_production) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id",
		form.Name,
		form.Goal,
		form.Details,
		time.Now(),
		form.Label,
		"NOT_STARTED",
		form.TimeEstimate,
		startedAtParsed,
		endedAtParsed,
		0,
		form.IsProduction).Scan(&strategyId)

	return strategyId, err
}

func (s StrategyModel) GetList() ([]Strategy, error) {
	rows, err := db.GetDB().Query("SELECT id, name, goal, details, created_at, label, status, time_estimate, started_at, ended_at, process, is_production FROM strategy ORDER BY status desc, updated_at desc, created_at desc")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var strategies []Strategy

	for rows.Next() {
		var strategy Strategy
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

	_, errorUpdate := db.GetDB().Exec("UPDATE strategy SET status=$1, updated_at=$2 WHERE id=$3", form.Status, time.Now(), form.Id)
	if errorUpdate != nil {
		return errorUpdate
	}

	_, errorSetHistory := db.GetDB().Exec("INSERT INTO strategy_changed_history(created_at, status, strategy_id) VALUES($1, $2, $3) RETURNING id", time.Now(), form.Status, form.Id)
	if errorSetHistory != nil {
		return errorSetHistory
	}

	return nil
}

func getPoromodoAverage(Id int) (utils.AverageScore, error) {
	rows, errPoromodo := db.GetDB().Query("SELECT satisfaction, productivity, interested, insight FROM poromodo WHERE id=$1", Id)
	if errPoromodo != nil {
		return utils.AverageScore{}, errPoromodo
	}

	defer rows.Close()

	var poromodos []utils.AverageScore

	for rows.Next() {
		var poromodo utils.AverageScore
		if err := rows.Scan(&poromodo.Satisfaction, &poromodo.Productivity, &poromodo.Interested, &poromodo.Insight); err != nil {
			return utils.AverageScore{}, err
		}
		poromodos = append(poromodos, poromodo)
	}
	if errRows := rows.Err(); errRows != nil {
		return utils.AverageScore{}, errRows
	}

	return utils.CalculateAverageScore(poromodos), nil
}

func (s StrategyModel) TriggerCompleted(form forms.StrategyCompletedForm) error {
	_, err := db.GetDB().Query("SELECT * FROM strategy WHERE id=$1", form.Id)
	if err != nil {
		return err
	}

	poromodoAverage, errorAverage := getPoromodoAverage(form.Id)
	if errorAverage != nil {
		return errorAverage
	}

	_, errorUpdate := db.GetDB().Exec("UPDATE strategy SET status=$1, updated_at=$2, satisfaction=$3, productivity=$4, interested=$5, insight=$6, completed_at=$7 WHERE id=$8", "COMPLETED", time.Now(), poromodoAverage.Satisfaction, poromodoAverage.Productivity, poromodoAverage.Interested, poromodoAverage.Insight, time.Now(), form.Id)
	if errorUpdate != nil {
		return errorUpdate
	}

	_, errorSetHistory := db.GetDB().Exec("INSERT INTO strategy_changed_history(created_at, status, strategy_id) VALUES($1, $2, $3) RETURNING id", time.Now(), "COMPLETED", form.Id)
	if errorSetHistory != nil {
		return errorSetHistory
	}

	return nil
}

func (s StrategyModel) GetStrategiesByStatus(status string) ([]ListStrategiesByStatus, error) {
	rows, err := db.GetDB().Query("SELECT id, name, process, goal, label FROM strategy WHERE status=$1", status)
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

func (s StrategyModel) GetDetailChangedHistory(id string) ([]ChangedHistory, error) {
	rows, err := db.GetDB().Query("SELECT created_at, status FROM strategy_changed_history WHERE strategy_id=$1 order by created_at desc", id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var changedHistories []ChangedHistory

	for rows.Next() {
		var changedHistory ChangedHistory
		if err := rows.Scan(&changedHistory.CreatedAt, &changedHistory.Status); err != nil {
			return changedHistories, err
		}
		changedHistories = append(changedHistories, changedHistory)
	}
	if err = rows.Err(); err != nil {
		return changedHistories, err
	}
	return changedHistories, nil
}
