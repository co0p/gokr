package gokr

type MeasurementStore interface {
	All() ([]Measurement, error)
	Create(Measurement) error
	Update(Measurement) error
	Delete(MeasurementId) error
}
