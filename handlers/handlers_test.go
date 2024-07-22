package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"test-ms-credit-card-validator/models"
	"testing"
)

var message = "Credit card number is valid"

func TestValidateCardHandler(t *testing.T) {
	tests := []struct {
		name         string
		cardNumber   string
		expectedCode int
		expectedBody models.ValidationResponse
	}{
		{
			name:         "Valid Visa Card",
			cardNumber:   "4111111111111111",
			expectedCode: http.StatusOK,
			expectedBody: models.ValidationResponse{
				IsValid:  true,
				Message:  message,
				CardType: "Visa",
			},
		},
		{
			name:         "Valid MasterCard Card",
			cardNumber:   "5555555555554444",
			expectedCode: http.StatusOK,
			expectedBody: models.ValidationResponse{
				IsValid:  true,
				Message:  message,
				CardType: "MasterCard",
			},
		},
		{
			name:         "Valid American Express Card",
			cardNumber:   "378282246310005",
			expectedCode: http.StatusOK,
			expectedBody: models.ValidationResponse{
				IsValid:  true,
				Message:  message,
				CardType: "American Express",
			},
		},
		{
			name:         "Valid Discover Card",
			cardNumber:   "6011111111111117",
			expectedCode: http.StatusOK,
			expectedBody: models.ValidationResponse{
				IsValid:  true,
				Message:  message,
				CardType: "Discover",
			},
		},
		{
			name:         "Valid Diners Club Card",
			cardNumber:   "30569309025904",
			expectedCode: http.StatusOK,
			expectedBody: models.ValidationResponse{
				IsValid:  true,
				Message:  message,
				CardType: "Diners Club",
			},
		},
		{
			name:         "Invalid Card",
			cardNumber:   "1234567890123456",
			expectedCode: http.StatusOK,
			expectedBody: models.ValidationResponse{
				IsValid:  false,
				Message:  "Credit card number is invalid",
				CardType: "Unknown",
			},
		},
		{
			name:         "Invalid Request Payload",
			cardNumber:   "",
			expectedCode: http.StatusBadRequest,
			expectedBody: models.ValidationResponse{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create the request body
			body, _ := json.Marshal(models.ValidationRequest{CardNumber: tt.cardNumber})
			req := httptest.NewRequest("POST", "/validate", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")

			// Create a ResponseRecorder to capture the response
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(ValidateCardHandler)

			// Serve the request
			handler.ServeHTTP(rr, req)

			// Check the status code
			if status := rr.Code; status != tt.expectedCode {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.expectedCode)
			}

			// Check the response body
			if tt.expectedCode == http.StatusOK {
				var response models.ValidationResponse

				if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
					t.Errorf("could not decode response: %v", err)
				}

				if response != tt.expectedBody {
					t.Errorf("handler returned unexpected body: got %v want %v",
						response, tt.expectedBody)
				}
			}
		})
	}
}
