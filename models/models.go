package models

type ValidationResponse struct {
	IsValid  bool   `json:"is_valid"`
	Message  string `json:"message"`
	CardType string `json:"card_type,omitempty"`
}

type ValidationRequest struct {
	CardNumber string `json:"card_number"`
}
