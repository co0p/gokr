package main

import (
	"crypto/subtle"
	"log"
	"net/http"
	"os"
	"text/template"
	"time"

	"gopkg.in/mgo.v2"
)

var tmp *template.Template
var session *mgo.Session

type CompanyStats struct {
	CreatedAt time.Time

	ProjectsBudgetPercentage uint
	ProjectsRunningCount     uint
	ProjectsMaintenanceCount uint
	PrototypesCreatedCount   uint
	IncidentCount            uint
	MaintenanceCost          int

	OfferCycletime                  uint
	OffersCreatedCount              uint
	OffersDeniedCount               uint
	InquieriesNewCustomerCount      uint
	InquieriesExistingCustomerCount uint
	InquieriesDeniedCount           uint

	BudgetPlanned        int
	BudgetCurrent        int
	OccupancyPercentage  uint
	MatchingGartnerCount uint
	MarketingCount       uint
}

func init() {
	tmp = template.Must(template.ParseFiles("index.html"))
}

func main() {

	// read environment vars
	port := envWithDefault("PORT", "8080")
	user := envWithDefault("GOKR_USERNAME", "user")
	pwd := envWithDefault("GOKR_PASSWORD", "user2")
	mongoURI := envWithDefault("GOKR_MONGO_PATH", "")

	log.Printf("loaded configuration from environment\n")

	s, err := mgo.Dial(mongoURI)
	session = s
	if err != nil {
		log.Fatalln("failed to connect to mongo: " + err.Error())
	}
	log.Printf("connected to mongodb\n")

	http.Handle("/", BasicAuth(handleHTTP, user, pwd))
	log.Printf("starting webserver at: %s\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalln("failed to start http server: " + err.Error())
	}
}

func handleHTTP(w http.ResponseWriter, r *http.Request) {
	s := session.Clone()
	defer s.Close()

	// aggregate result

	// create data

	current := CompanyStats{
		BudgetCurrent: 1,
	}

	previous := CompanyStats{
		BudgetCurrent: 1,
	}

	vm := struct {
		Current, Previous CompanyStats
	}{
		current,
		previous,
	}

	log.Printf("serving data: %v \n", vm)

	// serve template
	tmp.Execute(w, vm)
}

// BasicAuth wraps a given handlerFunc with basic auth protection
// copied shamelessly from https://stackoverflow.com/questions/21936332/idiomatic-way-of-requiring-http-basic-auth-in-go
func BasicAuth(handler http.HandlerFunc, username, password string) http.HandlerFunc {

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

func envWithDefault(key, defaultValue string) string {
	val := os.Getenv(key)
	if val == "" {
		val = defaultValue
	}
	return val
}
