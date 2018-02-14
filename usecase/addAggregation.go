package usecase

import (
	"log"

	"github.com/co0p/gokr"
)

type AddAggregation struct {
	Store gokr.AggregationStore
}

func (u *AddAggregation) Add(agg gokr.Aggregation) error {
	log.Printf("Storing aggregation: %s", agg)
	return u.Store.Save(agg)
}
