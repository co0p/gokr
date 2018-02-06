package gokr_test

import (
	"testing"
	"time"

	"github.com/co0p/gokr"

	"github.com/co0p/gokr/mock"
)

var measures = []gokr.Measurement{
	gokr.Measurement{Id: 12, Type: gokr.IncidentCount, CreatedAt: time.Now(), Value: 12},
	gokr.Measurement{Id: 22, Type: gokr.IncidentCount, CreatedAt: time.Now(), Value: 212},
}

func returnMeasurements() ([]gokr.Measurement, error) {
	return measures, nil
}

func TestMeasurementService(t *testing.T) {

	// create a mock
	var mockMeasurementStore mock.MeasurementStore
	mockMeasurementStore.GetAllOfTypeFn = returnMeasurements

	// construct testee
	srv := gokr.MeasurementService{Store: &mockMeasurementStore}

	t.Run("All", func(t *testing.T) {

		m, err := srv.All()
		if err != nil {
			t.Errorf("err should be nil, but got: %v", err)
		}

		if mockMeasurementStore.AllCount {
			t.Errorf("All() should been have called once, got: %v", mockMeasurementStore.AllCount)
		}
	})
}
