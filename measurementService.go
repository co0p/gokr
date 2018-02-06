package gokr

type MeasurementService struct {
	Store MeasurementStore
}

func (s MeasurementService) All() ([]Measurement, error) {
	return s.Store.All()
}

func (s MeasurementService) Create(m Measurement) error {
	return s.Store.Create(m)
}

func (s MeasurementService) Update(m Measurement) error {
	return s.Store.Update(m)
}

func (s MeasurementService) Delete(id MeasurementId) error {
	return s.Store.Delete(id)
}
