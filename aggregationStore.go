package gokr

type AggregationStore interface {
	All() ([]Aggregation, error)
	Latest() (Aggregation, error)
	Save(Aggregation) error
}
