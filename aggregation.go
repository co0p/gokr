package gokr

import "time"

type AggregationId int

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

	FromDate, ToDate time.Time
}
