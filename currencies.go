package fake

func Currency() string {
	return lookup(lang, "currencies", true)
}

func CurrencyCode() string {
	return lookup(lang, "currency_codes", true)
}
