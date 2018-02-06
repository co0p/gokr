package gokr

type AggregationStore interface {
	All() ([]Aggregation, error)
	Save(Aggregation) error
}
