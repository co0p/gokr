package mongo

import (
	"github.com/co0p/gokr"
	"gopkg.in/mgo.v2"
)

type AggregationStore struct {
	session *mgo.Session
}

func Connect(uri string) (AggregationStore, error) {
	session, err := mgo.Dial(uri)
	if err != nil {
		return AggregationStore{}, err
	}
	store := AggregationStore{session: session}
	return store, nil
}

func (s *AggregationStore) All() ([]gokr.Aggregation, error) {
	return nil, nil
}

func (s *AggregationStore) Save(gokr.Aggregation) error {
	return nil
}
