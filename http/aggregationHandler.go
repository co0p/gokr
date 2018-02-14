package http

import (
	"fmt"
	"log"
	"net/http"

	"github.com/co0p/gokr/usecase"
)

type AddAggregation struct {
	Usecase *usecase.AddAggregation
}

func (h AddAggregation) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("hi there")
	w.Write([]byte("hi there"))
}

type GetAggregation struct {
	Usecase *usecase.GetAggregation
}

func (h GetAggregation) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	docs, err := h.Usecase.All()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	msg := fmt.Sprintf("found %d docs", len(docs))
	w.Write([]byte(msg))

}
