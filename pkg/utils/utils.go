package utils

import "regexp"

func ValidateLuhn(cardNumber string) bool {
	var sum int
	alternate := false

	for i := len(cardNumber) - 1; i > -1; i-- {
		n := int(cardNumber[i] - '0')

		if alternate {
			n *= 2
			if n > 9 {
				n = (n % 10) + 1
			}
		}

		sum += n
		alternate = !alternate
	}

	return sum%10 == 0
}

func IsValidCardNumber(cardNumber string) bool {
	var re = regexp.MustCompile(`^[0-9]+$`)

	if !re.MatchString(cardNumber) {
		return false
	}

	return ValidateLuhn(cardNumber)
}

func GetCardType(cardNumber string) string {
	switch {
	case regexp.MustCompile(`^4[0-9]{12}(?:[0-9]{3})?$`).MatchString(cardNumber):
		return "Visa"

	case regexp.MustCompile(`^5[1-5][0-9]{14}$`).MatchString(cardNumber):
		return "MasterCard"

	case regexp.MustCompile(`^3[47][0-9]{13}$`).MatchString(cardNumber):
		return "American Express"

	case regexp.MustCompile(`^6(?:011|5[0-9]{2})[0-9]{12}$`).MatchString(cardNumber):
		return "Discover"

	case regexp.MustCompile(`^3(?:0[0-5]|[68][0-9])[0-9]{11}$`).MatchString(cardNumber):
		return "Diners Club"

	default:
		return "Unknown"
	}
}
