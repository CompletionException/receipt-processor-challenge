package api

import (
	"encoding/json"
	"log"
	"net/http"
	"receipt-processor-challenge/internal/business"
	"receipt-processor-challenge/internal/model"
	"receipt-processor-challenge/internal/storage"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// processReceiptHandler processes a new receipt and handle business logic.  Known ERROR always 400 per documentation
func processReceiptHandler(w http.ResponseWriter, r *http.Request, storage *storage.InMemoryStorage) {
	vars := mux.Vars(r)
	userIdStr := vars["userId"]

	if r.Method != http.MethodPost {
		//Invalid request method
		JSONError(w, "The receipt is invalid", http.StatusBadRequest)
		return
	}

	var receipt model.Receipt
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&receipt); err != nil {
		//Invalid request payload
		JSONError(w, "The receipt is invalid", http.StatusBadRequest)
		log.Print(err)
		return
	}

	// Generate a new UUID for the receipt
	receipt.ID = uuid.New()

	// Validate request parameters
	if (!business.Validate(receipt)){
		//Invalid request
		JSONError(w, "The receipt is invalid", http.StatusBadRequest)
		return
	}

	// Store user and calculate count
	if err := storage.SaveUserHistory(userIdStr); err != nil {
		//Failed to store user
		JSONError(w, "The receipt is invalid", http.StatusBadRequest)
		log.Print(err)
		return
	}
	// Retrieve the user count from in-memory storage
	userIdReceiptCount, found := storage.GetUserReceiptCount(userIdStr)
	if !found {
		//Receipt not found
		JSONError(w, "No receipt found for that id", http.StatusNotFound)
		return
	}

	// Calculate points based on the receipt
	receipt.Points = business.CalculatePoints(receipt, userIdReceiptCount)

	// Save receipt to in-memory storage
	if err := storage.SaveReceipt(receipt); err != nil {
		//Failed to save receipt
		JSONError(w, "The receipt is invalid", http.StatusBadRequest)
		log.Print(err)
		return
	}

	// Respond with a success message
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"id": receipt.ID.String(),
	}
	json.NewEncoder(w).Encode(response)
}

// getReceiptPointsHandler retrieves receipt data by ID. Known ERROR always 404 per documentation
func getReceiptPointsHandler(w http.ResponseWriter, r *http.Request, storage *storage.InMemoryStorage) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	// Parse ID from the URL
	id, err := uuid.Parse(idStr)
	if err != nil {
		//Invalid ID format
		JSONError(w, "No receipt found for that id", http.StatusNotFound)
		log.Print(err)
		return
	}

	// Retrieve the receipt from in-memory storage
	receipt, found := storage.GetReceipt(id)
	if !found {
		//Receipt not found
		JSONError(w, "No receipt found for that id", http.StatusNotFound)
		return
	}

	// Respond with the receipt points
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"points": receipt.Points,
	}
	json.NewEncoder(w).Encode(response)
}

// JSONEerror replaces http.Error. Return error as JSON to comply with documentation
func JSONError(w http.ResponseWriter, err interface{}, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(err)
}
