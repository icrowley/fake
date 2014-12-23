package test

import (
	"testing"

	"github.com/icrowley/fake"
)

func TestGeo(t *testing.T) {
	fake.SetLang("en")

	f := fake.Latitute()
	if f == 0 {
		t.Error("Latitude failed")
	}

	i := fake.LatitudeDegress()
	if i < -180 || i > 180 {
		t.Error("LatitudeDegress failed")
	}

	i = fake.LatitudeMinutes()
	if i < 0 || i >= 60 {
		t.Error("LatitudeMinutes failed")
	}

	i = fake.LatitudeSeconds()
	if i < 0 || i >= 60 {
		t.Error("LatitudeSeconds failed")
	}

	s := fake.LatitudeDirection()
	if s != "N" && s != "S" {
		t.Error("LatitudeDirection failed")
	}

	f = fake.Longitude()
	if f == 0 {
		t.Error("Longitude failed")
	}

	i = fake.LongitudeDegrees()
	if i < -180 || i > 180 {
		t.Error("LongitudeDegrees failed")
	}

	i = fake.LongitudeMinutes()
	if i < 0 || i >= 60 {
		t.Error("LongitudeMinutes failed")
	}

	i = fake.LongitudeSeconds()
	if i < 0 || i >= 60 {
		t.Error("LongitudeSeconds failed")
	}

	s = fake.LongitudeDirection()
	if s != "W" && s != "E" {
		t.Error("LongitudeDirection failed")
	}
}
