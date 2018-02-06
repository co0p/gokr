package gokr

import (
	"fmt"
)

type AggregationService struct {
	AggregationStore AggregationStore
	MeasurementStore MeasurementStore
}

func (s AggregationService) All() ([]Aggregation, error) {
	/* measures, err := s.measurementStore.All()
	if err != nil {
		return nil, fmt.Errorf("failed to get measurements: %s", err.Error())
	}

	measurementTypes := getMeasurementTypes(measures)

	for k, v := range measurementTypes {

	}
	value, err := applyAggregation(measures, v)

	if err != nil {
		return Aggregation{}, fmt.Errorf("failed to aggregate measurements: %s", err.Error())
	}

	agg := Aggregation{MeasurmentType: mType, FromDate: time.Unix(0, 0), ToDate: time.Now(), Value: value}
	return agg, nil

	*/
	return s.AggregationStore.All()
}

func applyAggregation(measures []Measurement, mType MeasurementType) (int, error) {
	if len(measures) == 0 {
		return 0, nil
	}

	switch mType {
	case IncidentCount, PrototypesCompletedCount:
		return sum(measures), nil
	case OccupancyPercentage:
		return avg(measures), nil
	default:
		return 0, fmt.Errorf("failed to aggregate for type: %v. No aggegate function found", mType)
	}
}

func getMeasurementTypes(measurements []Measurement) []MeasurementType {

	m := make(map[MeasurementType]bool)
	for _, v := range measurements {
		m[v.Type] = true
	}

	keys := make([]MeasurementType, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func sum(measures []Measurement) int {
	value := 0
	for _, v := range measures {
		value += v.Value
	}
	return value
}

func avg(measures []Measurement) int {
	value := 0
	for _, v := range measures {
		value += v.Value
	}
	return value / len(measures)
}
