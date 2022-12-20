package utils

const (
	// USD is the currency code for US Dollar
	USD = "USD"
	// EUR is the currency code for Euro
	EUR = "EUR"
	// AUD is the currency code for Australian Dollar
	AUD = "AUD"
	// SGD is the currency code for Singapore Dollar
	SGD = "SGD"
	// VND is the currency code for Vietnamese
	VND = "VND"
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, AUD, SGD, VND:
		return true
	default:
		return false
	}
}
