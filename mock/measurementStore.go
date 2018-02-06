package mock

import (
	"github.com/co0p/gokr"
)

type MeasurementStore struct {
	AllCount, CreateCalled, UpdateCalled, DeleteCalled bool

	AllFn    func() ([]gokr.Measurement, error)
	CreateFn func(gokr.Measurement) error
	UpdateFn func(gokr.Measurement) error
	DeleteFn func(gokr.MeasurementId) error
}

func (s *MeasurementStore) All() ([]gokr.Measurement, error) {
	s.AllCount = true
	return s.AllFn()
}

func (s *MeasurementStore) Create(m gokr.Measurement) error {
	s.CreateCalled = true
	return s.CreateFn(m)
}
func (s *MeasurementStore) Update(m gokr.Measurement) error {
	s.UpdateCalled = true
	return s.UpdateFn(m)
}
func (s *MeasurementStore) Delete(id gokr.MeasurementId) error {
	s.DeleteCalled = true
	return s.DeleteFn(id)
}
