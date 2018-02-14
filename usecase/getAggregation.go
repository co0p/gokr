package usecase

import (
	"log"

	"github.com/co0p/gokr"
)

type GetAggregation struct {
	Store gokr.AggregationStore
}

func (u *GetAggregation) All() ([]gokr.Aggregation, error) {
	log.Printf("get  all aggregations")
	return u.Store.All()
}

func (u *GetAggregation) Latest() (gokr.Aggregation, error) {
	log.Printf("get latest aggregation")
	return u.Store.Latest()
}
