package utils

import "strings"

func IdentifyCard(cardNumber string) string {
	if strings.HasPrefix(cardNumber, "4") {
		return "Visa"
	} else if strings.HasPrefix(cardNumber, "5") {

		firstTwoDigits := cardNumber[:2]
		if firstTwoDigits >= "51" && firstTwoDigits <= "55" {
			return "Mastercard"
		}

		firstFourDigits := cardNumber[:4]
		if firstFourDigits >= "2221" && firstFourDigits <= "2720" {
			return "Mastercard"
		}
	} else if strings.HasPrefix(cardNumber, "34") || strings.HasPrefix(cardNumber, "37") {
		return "American Express"
	} else if strings.HasPrefix(cardNumber, "6011") || (cardNumber[:3] >= "644" && cardNumber[:3] <= "649") || strings.HasPrefix(cardNumber, "65") {
		return "Discover"
	} else if cardNumber[:4] >= "3528" && cardNumber[:4] <= "3589" {
		return "JCB"
	} else if (cardNumber[:3] >= "300" && cardNumber[:3] <= "305") || strings.HasPrefix(cardNumber, "36") || (cardNumber[:2] >= "38" && cardNumber[:2] <= "39") {
		return "Diners Club"
	}
	return "Unknown"

}
