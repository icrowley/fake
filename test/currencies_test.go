package test

import (
	"testing"

	"github.com/icrowley/fake"
)

func TestCurrencies(t *testing.T) {
	fake.SetLang("en")

	v := fake.Currency()
	if v == "" {
		t.Error("Currency failed")
	}

	v = fake.CurrencyCode()
	if v == "" {
		t.Error("CurrencyCode failed")
	}
}
