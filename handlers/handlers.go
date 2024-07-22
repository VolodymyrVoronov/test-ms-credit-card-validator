package handlers

import (
	"encoding/json"
	"net/http"
	"test-ms-credit-card-validator/models"
	"test-ms-credit-card-validator/pkg/utils"
)

func ValidateCardHandler(w http.ResponseWriter, r *http.Request) {
	var req models.ValidationRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	cardNumber := req.CardNumber
	if cardNumber == "" {
		http.Error(w, "Missing card_number parameter", http.StatusBadRequest)
		return
	}

	isValid := utils.IsValidCardNumber(cardNumber)
	cardType := utils.GetCardType(cardNumber)

	response := models.ValidationResponse{
		IsValid:  isValid,
		CardType: cardType,
	}

	if !isValid {
		response.Message = "Credit card number is invalid"
	} else {
		response.Message = "Credit card number is valid"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
