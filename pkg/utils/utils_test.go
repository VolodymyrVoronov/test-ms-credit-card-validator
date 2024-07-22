package utils

import (
	"testing"
)

func TestValidateLun(t *testing.T) {
	tests := []struct {
		name       string
		cardNumber string
		expected   bool
	}{
		{"Valid Visa", "4111111111111111", true},
		{"Valid MasterCard", "5555555555554444", true},
		{"Valid American Express", "378282246310005", true},
		{"Valid Discover", "6011111111111117", true},
		{"Valid Diners Club", "30569309025904", true},
		{"Invalid Card", "1234567890123456", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ValidateLuhn(tt.cardNumber)

			if result != tt.expected {
				t.Errorf("ValidateLun(%s) = %v; want %v", tt.cardNumber, result, tt.expected)
			}
		})
	}
}

func TestIsValidCardNumber(t *testing.T) {
	tests := []struct {
		name       string
		cardNumber string
		expected   bool
	}{
		{"Valid Visa", "4111111111111111", true},
		{"Valid MasterCard", "5555555555554444", true},
		{"Valid American Express", "378282246310005", true},
		{"Valid Discover", "6011111111111117", true},
		{"Valid Diners Club", "30569309025904", true},
		{"Invalid Card", "1234567890123456", false},
		{"Invalid Characters", "4111-1111-1111-1111", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsValidCardNumber(tt.cardNumber)

			if result != tt.expected {
				t.Errorf("IsValidCardNumber(%s) = %v; want %v", tt.cardNumber, result, tt.expected)
			}
		})
	}
}

func TestGetCardType(t *testing.T) {
	tests := []struct {
		name       string
		cardNumber string
		expected   string
	}{
		{"Visa", "4111111111111111", "Visa"},
		{"MasterCard", "5105105105105100", "MasterCard"},
		{"American Express", "378282246310005", "American Express"},
		{"Discover", "6011111111111117", "Discover"},
		{"Diners Club", "30569309025904", "Diners Club"},
		{"Unknown", "1234567890123456", "Unknown"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetCardType(tt.cardNumber)

			if result != tt.expected {
				t.Errorf("GetCardType(%s) = %v; want %v", tt.cardNumber, result, tt.expected)
			}
		})
	}
}
