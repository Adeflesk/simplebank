package util

// constants for all supported currencies
const (
	USD = "USD"
	EUR = "EUR"
	GBP = "GBP"
	JPY = "JPY"
	CNY = "CNY"
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, GBP, JPY, CNY:
		return true
	}
	return false
}
