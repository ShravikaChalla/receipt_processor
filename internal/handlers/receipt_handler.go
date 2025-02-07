package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"receipt_processor/internal/models"
	"receipt_processor/internal/services"

	"github.com/gorilla/mux"
)

var receiptService = services.NewReceiptService()

func ProcessReceiptHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Begin - ProcessReceiptHandler...")
	var receipt models.Receipt
	if err := json.NewDecoder(r.Body).Decode(&receipt); err != nil {
		http.Error(w, "Invalid receipt format", http.StatusBadRequest)
		return
	}
	id := receiptService.ProcessReceipt(receipt)
	log.Println("id = ", id)
	log.Println("End - ProcessReceiptHandler...")
	json.NewEncoder(w).Encode(map[string]string{"id": id})
}

func GetPointsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Begin - GetPointsHandler...")
	vars := mux.Vars(r)
	id := vars["id"]
	points, exists := receiptService.GetPoints(id)
	if !exists {
		http.Error(w, "No receipt found for that ID.", http.StatusNotFound)
		return
	}
	log.Println("points = ", points)
	log.Println("End - GetPointsHandler...")
	json.NewEncoder(w).Encode(map[string]int{"points": points})
}
