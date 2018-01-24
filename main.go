package main

import (
	"crypto/subtle"
	"html/template"
	"log"
	"net/http"
	"os"

	"gopkg.in/mgo.v2"
)

const realm = "OKRs done in go"

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
	mongoURI := envWithDefault("GOKR_MONGO_PATH", "mongodb://gokr:golangokrs@ds113738.mlab.com:13738/gokr")

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

	vm := struct {
		Title    string
		Headline string
	}{
		"gokr",
		"OKRs done in go",
	}

	// serve template
	tmp.Execute(w, vm)
}

// BasicAuth wraps a given handlerFunc with basic auth protection
// copied shamelessly from https://stackoverflow.com/questions/21936332/idiomatic-way-of-requiring-http-basic-auth-in-go
func BasicAuth(handler http.HandlerFunc, username, password string) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		user, pass, ok := r.BasicAuth()

		if !ok || subtle.ConstantTimeCompare([]byte(user), []byte(username)) != 1 || subtle.ConstantTimeCompare([]byte(pass), []byte(password)) != 1 {
			w.Header().Set("WWW-Authenticate", `Basic realm="`+realm+`"`)
			w.WriteHeader(401)
			w.Write([]byte("Unauthorised.\n"))
			log.Printf("rejected access for user=%s pwd=%s\n", user, pass)
			return
		}

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
