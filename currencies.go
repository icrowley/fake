package fake

// Currency generates currency name
func Currency() string {
	return f.Currency()
}

// CurrencyCode generates currency code
func CurrencyCode() string {
	return f.CurrencyCode()
}

// Currency generates currency name
func (f *Faker) Currency() string {
	return f.lookup(f.lang, "currencies", true)
}

// CurrencyCode generates currency code
func (f *Faker) CurrencyCode() string {
	return f.lookup(f.lang, "currency_codes", true)
}
