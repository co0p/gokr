package main

import (
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/co0p/gokr"

	gokrhttp "github.com/co0p/gokr/http"
	"github.com/co0p/gokr/mongo"
)

var tmp *template.Template

func init() {
	tmp = template.Must(template.ParseFiles("templates/index.html"))
}

func main() {

	// read environment vars
	port := envWithDefault("PORT", "8080")
	//user := envWithDefault("GOKR_USERNAME", "user")
	//pwd := envWithDefault("GOKR_PASSWORD", "user2")
	mongoURI := envWithDefault("GOKR_MONGO_PATH", "")
	log.Printf("loaded configuration from environment\n")

	// connect to mongo
	store, err := mongo.Connect(mongoURI)
	if err != nil {
		log.Fatalln("failed to connect to mongo: " + err.Error())
	}
	log.Printf("connected to mongodb\n")

	// assemble dependencies
	aggregationService := gokr.AggregationService{AggregationStore: &store}
	aggregationHandler := gokrhttp.AggregationHandler{AggregationService: &aggregationService}

	mux := http.NewServeMux()
	mux.HandleFunc("/", aggregationHandler.Handle())

	log.Printf("starting webserver at: %s\n", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
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
