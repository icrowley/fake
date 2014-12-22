package fake

func Currency() string {
	return lookup("currency", "currencies", lang, true)
}

func CurrencyCode() string {
	return lookup("currency", "currency_codes", lang, true)
}
