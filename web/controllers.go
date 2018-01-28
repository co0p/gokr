package web

import (
	"crypto/subtle"
	"encoding/json"
	"log"
	"net/http"
	"text/template"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/co0p/gokr/events"
	"gopkg.in/mgo.v2"
)

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

	eventTypes := []events.EventType{events.IncidentCount, events.MaintenanceCost, events.ProjectsMaintenanceCount, events.ProjectsRunningCount, events.PrototypesCreatedCount}

	aggregatedStats := AggregatedStats{
		IncidentCount:            1,
		MaintenanceCost:          10000,
		PrototypesCompletedCount: 0,
	}

	vm := struct {
		EventTypes []events.EventType
		Stats      AggregatedStats
	}{
		EventTypes: eventTypes,
		Stats:      aggregatedStats,
	}

	// serve template
	h.template.Execute(w, vm)
}

func (h Handlers) EventsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		log.Printf("unsupported method=%s", r.Method)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	var e events.Event
	if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// create a new event id
	e.ID = bson.NewObjectId()
	e.Created = time.Now()

	// insert it into the database
	if err := h.session.DB("gokr").C("events").Insert(&e); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// WithBasicAuth wraps a given handlerFunc with basic auth protection
// copied shamelessly from https://stackoverflow.com/questions/21936332/idiomatic-way-of-requiring-http-basic-auth-in-go
func WithBasicAuth(handler http.HandlerFunc, username, password string) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		user, pass, ok := r.BasicAuth()

		if !ok || subtle.ConstantTimeCompare([]byte(user), []byte(username)) != 1 || subtle.ConstantTimeCompare([]byte(pass), []byte(password)) != 1 {
			w.Header().Set("WWW-Authenticate", `Basic realm="gokrs"`)
			w.WriteHeader(401)
			w.Write([]byte("Unauthorised.\n"))
			log.Printf("rejected access for user=%s pwd=%s\n", user, pass)
			return
		}

		log.Printf("granting access for user=%s pwd=%s\n", user, pass)

		handler(w, r)
	}
}
