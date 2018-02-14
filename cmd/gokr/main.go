package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/co0p/gokr/usecase"
	"github.com/gorilla/mux"

	httpHandler "github.com/co0p/gokr/http"
	"github.com/co0p/gokr/mongo"
)

func main() {

	// read environment vars
	port := envWithDefault("PORT", "8080")
	mongoURI := envWithDefault("GOKR_MONGO_PATH", "")
	log.Printf("loaded configuration from environment\n")

	// connect to mongo
	store, err := mongo.Connect(mongoURI)
	defer store.Close()
	if err != nil {
		log.Fatalln("failed to connect to mongo: " + err.Error())
	}
	log.Printf("connected to mongodb\n")

	// assemble dependencies
	getUsecase := usecase.GetAggregation{Store: &store}
	addUsecase := usecase.AddAggregation{Store: &store}

	getHandler := httpHandler.GetAggregation{Usecase: &getUsecase}
	addHandler := httpHandler.AddAggregation{Usecase: &addUsecase}

	r := mux.NewRouter()
	r.HandleFunc("/api/aggregations", getHandler.ServeHTTP).Methods("GET")
	r.HandleFunc("/api/aggregations", addHandler.ServeHTTP).Methods("POST")
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	srv := &http.Server{
		Handler: r,
		Addr:    ":" + port,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Printf("starting webserver at: %s\n", port)
	if err := srv.ListenAndServe(); err != nil {
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
