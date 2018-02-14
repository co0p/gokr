package mongo

import (
	"fmt"
	"log"
	"time"

	"github.com/co0p/gokr"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	AGGREGATION_COLLECTION = "aggregation"
	DATABASE_NAME          = "gokr"
)

type AggregationDocument struct {
	ID *bson.ObjectId
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

type AggregationStore struct {
	session  *mgo.Session
	database *mgo.Database
}

func Connect(uri string) (AggregationStore, error) {
	session, err := mgo.Dial(uri)
	if err != nil {
		return AggregationStore{}, err
	}
	database := session.DB(DATABASE_NAME)
	store := AggregationStore{session: session, database: database}
	return store, nil
}

func (s *AggregationStore) Close() {
	log.Println("closing mongo session")
	s.session.Close()
}

func (s *AggregationStore) All() ([]gokr.Aggregation, error) {
	collection := s.database.C(AGGREGATION_COLLECTION).With(s.session.Copy())
	var documents []AggregationDocument
	if err := collection.Find(bson.M{}).Sort("-createdAt").All(&documents); err != nil {
		return nil, fmt.Errorf("failed fetching all aggregations: %s", err.Error())
	}

	log.Printf("found %d documents", len(documents))
	aggregations := []gokr.Aggregation{}
	for _, v := range documents {
		aggregations = append(aggregations, gokr.Aggregation{
			ID:                              gokr.AggregationId(v.ID.String()),
			IncidentCount:                   v.IncidentCount,
			ProjectsRunningCount:            v.ProjectsRunningCount,
			ProjectsMaintenanceCount:        v.ProjectsMaintenanceCount,
			PrototypesCompletedCount:        v.PrototypesCompletedCount,
			MaintenanceCost:                 v.MaintenanceCost,
			ProjectsBudgetPercentage:        v.ProjectsBudgetPercentage,
			OffersCreatedCount:              v.OffersCreatedCount,
			OffersDeniedCount:               v.OffersDeniedCount,
			OfferCycletime:                  v.OfferCycletime,
			InquieriesNewCustomerCount:      v.InquieriesNewCustomerCount,
			InquieriesExistingCustomerCount: v.InquieriesExistingCustomerCount,
			InquieriesDeniedCount:           v.InquieriesDeniedCount,
			BudgetPlanned:                   v.BudgetPlanned,
			BudgetCurrent:                   v.BudgetCurrent,
			OccupancyPercentage:             v.OccupancyPercentage,
			GartnerHypeCount:                v.GartnerHypeCount,
			MarketingActivityCount:          v.MarketingActivityCount,
			CreatedAt:                       v.CreatedAt,
		})
	}
	return aggregations, nil
}

func (s *AggregationStore) Save(gokr.Aggregation) error {
	return nil
}

func (s *AggregationStore) Latest() (gokr.Aggregation, error) {
	collection := s.database.C(AGGREGATION_COLLECTION)
	var document AggregationDocument
	if err := collection.Find(bson.M{}).Sort("-createdAt").One(&document); err != nil {
		return gokr.Aggregation{}, fmt.Errorf("failed fetching latest aggregation: %s", err.Error())
	}

	aggregation := gokr.Aggregation{
		ID:                              gokr.AggregationId(document.ID.String()),
		IncidentCount:                   document.IncidentCount,
		ProjectsRunningCount:            document.ProjectsRunningCount,
		ProjectsMaintenanceCount:        document.ProjectsMaintenanceCount,
		PrototypesCompletedCount:        document.PrototypesCompletedCount,
		MaintenanceCost:                 document.MaintenanceCost,
		ProjectsBudgetPercentage:        document.ProjectsBudgetPercentage,
		OffersCreatedCount:              document.OffersCreatedCount,
		OffersDeniedCount:               document.OffersDeniedCount,
		OfferCycletime:                  document.OfferCycletime,
		InquieriesNewCustomerCount:      document.InquieriesNewCustomerCount,
		InquieriesExistingCustomerCount: document.InquieriesExistingCustomerCount,
		InquieriesDeniedCount:           document.InquieriesDeniedCount,
		BudgetPlanned:                   document.BudgetPlanned,
		BudgetCurrent:                   document.BudgetCurrent,
		OccupancyPercentage:             document.OccupancyPercentage,
		GartnerHypeCount:                document.GartnerHypeCount,
		MarketingActivityCount:          document.MarketingActivityCount,
		CreatedAt:                       document.CreatedAt,
	}
	return aggregation, nil
}
