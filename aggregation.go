package gokr

import "time"

type AggregationType string

const (
	Sum          AggregationType = "sum"
	TotalSum                     = "totalSum"
	Average                      = "average"
	TotalAverage                 = "totalAverage"
)

type Aggregation struct {
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

	FromDate, ToDate time.Time
}
