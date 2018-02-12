package usecase

import (
	"github.com/co0p/gokr"
)

type GetAggregation struct {
	Store gokr.AggregationStore
}

func (u *GetAggregation) All() ([]gokr.Aggregation, error) {
	return nil, nil
}

func (u *GetAggregation) Latest() (gokr.Aggregation, error) {
	return gokr.Aggregation{}, nil
}
