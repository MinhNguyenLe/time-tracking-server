package models

import (
	"github.com/MinhNguyenLe/time-tracking-server/db"
	"github.com/MinhNguyenLe/time-tracking-server/forms"
)

type StrategyModel struct{}

func (p StrategyModel) Insert(form forms.InsertStrategyForm) (strategyId int64, err error) {
	err = db.GetDB().QueryRow(
		"INSERT INTO public.strategy(name, goal, details, created_at, label, status, time_estimate, started_at, ended_at) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id",
		form.Name,
		form.Goal,
		form.Details,
		form.CreatedAt,
		form.Label,
		form.Status,
		form.TimeEstimate,
		form.StartedAt,
		form.EndedAt).Scan(&strategyId)

	return strategyId, err
}
