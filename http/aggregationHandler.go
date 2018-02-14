package http

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/co0p/gokr"
	"github.com/co0p/gokr/usecase"
)

type AggregationDto struct {
	ID string `json:"id"`

	IncidentCount                   int `json:"incidents"`
	ProjectsRunningCount            int `json:"projectsRunning"`
	ProjectsMaintenanceCount        int `json:"projectsMaintenance"`
	PrototypesCompletedCount        int `json:"prototypesCompleted"`
	MaintenanceCost                 int `json:"maintenanceCost"`
	ProjectsBudgetPercentage        int `json:"projectsBudgetPercentage"`
	OffersCreatedCount              int `json:"offersCreated"`
	OffersDeniedCount               int `json:"offersDenied"`
	OfferCycletime                  int `json:"offerCycletime"`
	InquieriesNewCustomerCount      int `json:"inquieriesNewCustomer"`
	InquieriesExistingCustomerCount int `json:"inquieriesExistingCustomer"`
	InquieriesDeniedCount           int `json:"inquieriesDenied"`
	BudgetPlanned                   int `json:"budgetPlanned"`
	BudgetCurrent                   int `json:"budgetCurrent"`
	OccupancyPercentage             int `json:"occupancyPercentage"`
	GartnerHypeCount                int `json:"gartnerHype"`
	MarketingActivityCount          int `json:"marketingActivity"`

	CreatedAt time.Time `json:"dateCreated"`
}

func toAggregationDto(a gokr.Aggregation) AggregationDto {
	return AggregationDto{
		ID:                              string(a.ID),
		IncidentCount:                   a.IncidentCount,
		ProjectsRunningCount:            a.ProjectsRunningCount,
		ProjectsMaintenanceCount:        a.ProjectsMaintenanceCount,
		PrototypesCompletedCount:        a.PrototypesCompletedCount,
		MaintenanceCost:                 a.MaintenanceCost,
		ProjectsBudgetPercentage:        a.ProjectsBudgetPercentage,
		OffersCreatedCount:              a.OffersCreatedCount,
		OffersDeniedCount:               a.OffersDeniedCount,
		OfferCycletime:                  a.OfferCycletime,
		InquieriesNewCustomerCount:      a.InquieriesNewCustomerCount,
		InquieriesExistingCustomerCount: a.InquieriesExistingCustomerCount,
		InquieriesDeniedCount:           a.InquieriesDeniedCount,
		BudgetPlanned:                   a.BudgetPlanned,
		BudgetCurrent:                   a.BudgetCurrent,
		OccupancyPercentage:             a.OccupancyPercentage,
		GartnerHypeCount:                a.GartnerHypeCount,
		MarketingActivityCount:          a.MarketingActivityCount,
		CreatedAt:                       a.CreatedAt,
	}
}

type AddAggregation struct {
	Usecase *usecase.AddAggregation
}

func (h AddAggregation) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("hi there")
	w.Write([]byte("hi there"))
}

type GetAggregation struct {
	Usecase *usecase.GetAggregation
}

func (h GetAggregation) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	docs, err := h.Usecase.All()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	dtos := []AggregationDto{}
	for _, v := range docs {
		dtos = append(dtos, toAggregationDto(v))
	}

	marshalled, err := json.Marshal(dtos)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(marshalled)
}
