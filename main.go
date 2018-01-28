package main

import (
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/co0p/gokr/web"

	"gopkg.in/mgo.v2"
)

var tmp *template.Template
var session *mgo.Session

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

	// connect to mongo
	session, err := mgo.Dial(mongoURI)
	if err != nil {
		log.Fatalln("failed to connect to mongo: " + err.Error())
	}
	log.Printf("connected to mongodb\n")

	// register handlers
	h := web.NewHandlers(tmp, session)

	http.Handle("/", web.WithBasicAuth(h.RootHandler, user, pwd))
	http.Handle("/api/events", web.WithBasicAuth(h.EventsHandler, user, pwd))

	log.Printf("starting webserver at: %s\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalln("failed to start http server: " + err.Error())
	}
}

func envWithDefault(key, defaultValue string) string {
	val := os.Getenv(key)
	if val == "" {
		val = defaultValue
	}
	return val
}
