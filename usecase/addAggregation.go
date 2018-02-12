package usecase

import (
	"github.com/co0p/gokr"
)

type AddAggregation struct {
	Store gokr.AggregationStore
}

func (u *AddAggregation) Add(m gokr.Aggregation) error {
	return nil
}
