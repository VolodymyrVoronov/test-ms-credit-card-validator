package models

import (
	"encoding/json"
	"testing"
)

func TestValidationRequestSerialization(t *testing.T) {
	expectedCardNumber := "4111111111111111"
	req := ValidationRequest{CardNumber: expectedCardNumber}

	// Test JSON marshaling
	data, err := json.Marshal(req)
	if err != nil {
		t.Errorf("Failed to marshal ValidationRequest: %v", err)
	}

	var unmarshaledReq ValidationRequest
	if err := json.Unmarshal(data, &unmarshaledReq); err != nil {
		t.Errorf("Failed to unmarshal ValidationRequest: %v", err)
	}

	if unmarshaledReq.CardNumber != expectedCardNumber {
		t.Errorf("Unmarshaled CardNumber = %v; want %v", unmarshaledReq.CardNumber, expectedCardNumber)
	}
}

func TestValidationResponseSerialization(t *testing.T) {
	expectedResponse := ValidationResponse{
		IsValid:  true,
		Message:  "Credit card number is valid",
		CardType: "Visa",
	}

	// Test JSON marshaling
	data, err := json.Marshal(expectedResponse)
	if err != nil {
		t.Errorf("Failed to marshal ValidationResponse: %v", err)
	}

	var unmarshaledResp ValidationResponse
	if err := json.Unmarshal(data, &unmarshaledResp); err != nil {
		t.Errorf("Failed to unmarshal ValidationResponse: %v", err)
	}

	if unmarshaledResp.IsValid != expectedResponse.IsValid {
		t.Errorf("Unmarshaled IsValid = %v; want %v", unmarshaledResp.IsValid, expectedResponse.IsValid)
	}
	if unmarshaledResp.Message != expectedResponse.Message {
		t.Errorf("Unmarshaled Message = %v; want %v", unmarshaledResp.Message, expectedResponse.Message)
	}
	if unmarshaledResp.CardType != expectedResponse.CardType {
		t.Errorf("Unmarshaled CardType = %v; want %v", unmarshaledResp.CardType, expectedResponse.CardType)
	}
}
