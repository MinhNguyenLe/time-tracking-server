package models

import (
	"github.com/MinhNguyenLe/time-tracking-server/db"
	"github.com/MinhNguyenLe/time-tracking-server/forms"
	"github.com/lib/pq"
)

type InsertPoromodoForm struct {
}

type PoromodoModel struct{}

func (p PoromodoModel) Insert(form forms.InsertPoromodoForm) (poromodoId int64, err error) {
	err = db.GetDB().QueryRow("INSERT INTO public.poromodo(goal, details) VALUES($1, $2) RETURNING id", form.Goal, pq.Array(form.Details)).Scan(&poromodoId)

	return poromodoId, err
}
