package gokr

import "time"

type MeasurementId int

type MeasurementType string

const (
	IncidentCount                   MeasurementType = "IncidentCount"
	ProjectsRunningCount                            = "ProjectsRunningCount"
	ProjectsMaintenanceCount                        = "ProjectsMaintenanceCount"
	PrototypesCompletedCount                        = "PrototypesCompletedCount"
	MaintenanceCost                                 = "MaintenanceCost"
	ProjectsBudgetPercentage                        = "ProjectsBudgetPercentage"
	OffersCreatedCount                              = "OffersCreatedCount"
	OffersDeniedCount                               = "OffersDeniedCount"
	OfferCycletime                                  = "OfferCycletime"
	InquieriesNewCustomerCount                      = "InquieriesNewCustomerCount"
	InquieriesExistingCustomerCount                 = "InquieriesExistingCustomerCount"
	InquieriesDeniedCount                           = "InquieriesDeniedCount"
	BudgetPlanned                                   = "BudgetPlanned"
	BudgetCurrent                                   = "BudgetCurrent"
	OccupancyPercentage                             = "OccupancyPercentage"
	GartnerHypeCount                                = "GartnerHypeCount"
	MarketingActivityCount                          = "MarketingActivityCount"
)

type Measurement struct {
	ID        MeasurementId
	Type      MeasurementType
	CreatedAt time.Time
	Value     int
	Note      string
}
