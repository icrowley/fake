package fake_test

import (
	"testing"

	"github.com/icrowley/fake"
)

func TestCreditCards(t *testing.T) {
	fake.SetLang("en")

	v := fake.CreditCardType()
	if v == "" {
		t.Error("CreditCardType failed")
	}

	v = fake.CreditCardNum("")
	if v == "" {
		t.Error("CreditCardNum failed")
	}

	v = fake.CreditCardNum("visa")
	if v == "" {
		t.Error("CreditCardNum failed")
	}
}
