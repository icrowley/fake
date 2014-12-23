package test

import (
	"testing"

	"github.com/icrowley/fake"
)

func TestCurrencies(t *testing.T) {
	langs := []string{"en", "ru"}
	for _, lang := range langs {
		fake.SetLang(lang)

		v := fake.Currency()
		if v == "" {
			t.Errorf("Currency failed with lang %s", lang)
		}

		v = fake.CurrencyCode()
		if v == "" {
			t.Errorf("CurrencyCode failed with lang %s", lang)
		}
	}
}
