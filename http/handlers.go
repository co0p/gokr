package http

import (
	"net/http"
	"text/template"

	"gopkg.in/mgo.v2"
)

type Measurement struct {
	ID int
}

type Handlers struct {
	template *template.Template
	session  *mgo.Session
}

func NewHandlers(template *template.Template, session *mgo.Session) Handlers {
	return Handlers{
		session:  session,
		template: template,
	}
}

func (h Handlers) RootHandler(w http.ResponseWriter, r *http.Request) {
	s := h.session.Copy()
	defer s.Close()

	// aggregate result

	// create data

	type AggregatedStats struct {
		IncidentCount            uint
		ProjectsRunningCount     uint
		ProjectsMaintenanceCount uint
		MaintenanceCost          uint
		PrototypesCompletedCount uint
	}

	aggregatedStats := AggregatedStats{
		IncidentCount:            1,
		MaintenanceCost:          10000,
		PrototypesCompletedCount: 0,
	}

	vm := struct {
		EventTypes []string
		Stats      AggregatedStats
	}{
		Stats: aggregatedStats,
	}

	// serve template
	h.template.Execute(w, vm)
}

func (h Handlers) EventsHandler(w http.ResponseWriter, r *http.Request) {
	/*
		if r.Method != "POST" {
			log.Printf("unsupported method=%s", r.Method)
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		var cmd Measurement
		if err := json.NewDecoder(r.Body).Decode(&cmd); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// create a new event id
		cmd.ID = bson.NewObjectId()

		// insert it into the database
		if err := h.session.DB("gokr").C("measurements").Insert(&cmd); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusCreated)
	*/
}
