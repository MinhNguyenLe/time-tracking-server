package models

import (
	"time"

	"github.com/MinhNguyenLe/time-tracking-server/db"
	"github.com/MinhNguyenLe/time-tracking-server/forms"
)

type PoromodoModel struct{}

func (p PoromodoModel) Insert(form forms.InsertPoromodoForm) (poromodoId int64, err error) {
	err = db.GetDB().QueryRow("INSERT INTO poromodo(duration, strategy_id, satisfaction, productivity, interested, insight, created_at, goal) VALUES($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id", form.Duration, form.StrategyId, form.Satisfaction, form.Productivity, form.Interested, form.Insight, time.Now(), form.Goal).Scan(&poromodoId)

	_, errorUpdate := db.GetDB().Query("UPDATE strategy SET process=process + $1, updated_at=$2 WHERE id=$3", form.Duration, time.Now(), form.StrategyId)
	if errorUpdate != nil {
		return 0, errorUpdate
	}

	return poromodoId, err
}
