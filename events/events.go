package events

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type EventType string

const (
	ProjectsRunningCount     EventType = "ProjectsRunningCount"
	ProjectsMaintenanceCount           = "ProjectsMaintenanceCount"
	PrototypesCreatedCount             = "PrototypesCreatedCount"
	IncidentCount                      = "IncidentCount"
	MaintenanceCost                    = "MaintenanceCost"
)

type Event struct {
	ID      bson.ObjectId `json:"id"`
	Created time.Time     `json:"createdAt"`
	Type    EventType     `json:"eventType"`
	Value   int           `json:"value"`
}

type StatsAggregated struct {
	CreatedAt time.Time

	ProjectsRunningCount     uint `json:"projectsInDevelopment"`
	ProjectsMaintenanceCount uint `json:"projectsInMaintenance"`
	PrototypesCompletedCount uint `json:"prototypesCompleted"`
	IncidentCount            uint `json:"incidencts"`
	MaintenanceCost          int  `json:"maintenanceCost"`
}
