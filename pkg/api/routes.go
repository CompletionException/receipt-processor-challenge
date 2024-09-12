package api

import (
	"net/http"
	"receipt-processor-challenge/internal/storage"

	"github.com/gorilla/mux"
)

// SetupRouter sets up the routes and returns the router
func SetupRouter(s *storage.InMemoryStorage) *mux.Router {
	r := mux.NewRouter()
	InitRoutes(r, s)
	return r
}

// InitRoutes initializes the API routes
func InitRoutes(r *mux.Router, s *storage.InMemoryStorage) {
	r.HandleFunc("/receipts/process/{userId}", func(w http.ResponseWriter, r *http.Request) {
		processReceiptHandler(w, r, s)
	}).Methods("POST")
    r.HandleFunc("/receipts/{id}/points", func(w http.ResponseWriter, r *http.Request) {
		getReceiptPointsHandler(w, r, s)
	}).Methods("GET")
}
