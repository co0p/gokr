package http

import (
	"net/http"

	"github.com/co0p/gokr"
)

type AggregationHandler struct {
	AggregationService *gokr.AggregationService
}

func (h *AggregationHandler) Handle() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hi there"))
	}
}
