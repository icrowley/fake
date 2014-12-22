package fake_test

import (
	"testing"

	"github.com/icrowley/fake"
)

func TestAddresses(t *testing.T) {
	fake.SetLang("en")

	v := fake.Continent()
	if v == "" {
		t.Error("Continent failed")
	}

	v = fake.Country()
	if v == "" {
		t.Error("Country failed")
	}

	v = fake.City()
	if v == "" {
		t.Error("City failed")
	}

	v = fake.State()
	if v == "" {
		t.Error("State failed")
	}

	v = fake.StateAbbrev()
	if v == "" {
		t.Error("StateAbbrev failed")
	}

	v = fake.Street()
	if v == "" {
		t.Error("Street failed")
	}

	v = fake.StreetAddress()
	if v == "" {
		t.Error("StreetAddress failed")
	}

	v = fake.Zip()
	if v == "" {
		t.Error("Zip failed")
	}

	v = fake.Phone()
	if v == "" {
		t.Error("Phone failed")
	}
}
