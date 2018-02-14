package gokr

import (
	"fmt"
	"time"
)

type AggregationId string

type Aggregation struct {
	ID AggregationId

	IncidentCount,
	ProjectsRunningCount,
	ProjectsMaintenanceCount,
	PrototypesCompletedCount,
	MaintenanceCost,
	ProjectsBudgetPercentage,
	OffersCreatedCount,
	OffersDeniedCount,
	OfferCycletime,
	InquieriesNewCustomerCount,
	InquieriesExistingCustomerCount,
	InquieriesDeniedCount,
	BudgetPlanned,
	BudgetCurrent,
	OccupancyPercentage,
	GartnerHypeCount,
	MarketingActivityCount int

	CreatedAt time.Time
}

func (a Aggregation) Stringer() string {
	return fmt.Sprintf("id: %s, created at %s", a.ID, a.CreatedAt)
}
