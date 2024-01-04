package models

import (
	"time"

	"github.com/MinhNguyenLe/time-tracking-server/db"
	"github.com/MinhNguyenLe/time-tracking-server/forms"
)

type PoromodoModel struct{}

func (p PoromodoModel) Insert(form forms.InsertPoromodoForm) (poromodoId int64, err error) {
	err = db.GetDB().QueryRow("INSERT INTO public.poromodo(duration, id_strategy, satisfaction, productivity, interested, insight, created_at) VALUES($1, $2, $3, $4, $5, $6, $7) RETURNING id", form.Duration, form.StrategyId, form.Satisfaction, form.Productivity, form.Interested, form.Insight, time.Now()).Scan(&poromodoId)

	return poromodoId, err
}
