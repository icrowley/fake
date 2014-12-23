package test

import (
	"testing"

	"github.com/icrowley/fake"
)

func TestGeneral(t *testing.T) {
	fake.SetLang("en")

	v := fake.Password(4, 10, true, true, true)
	if v == "" {
		t.Error("Password failed")
	}

	v = fake.SimplePassword()
	if v == "" {
		t.Error("SimplePassword failed")
	}

	v = fake.Color()
	if v == "" {
		t.Error("Color failed")
	}

	v = fake.HexColor()
	if v == "" {
		t.Error("HexColor failed")
	}

	v = fake.ShortHexColor()
	if v == "" {
		t.Error("ShortHexColor failed")
	}
}
