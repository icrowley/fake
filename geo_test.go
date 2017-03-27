package fake

import (
	"testing"
)

func TestGeo(t *testing.T) {
	for _, lang := range GetLangs() {
		SetLang(lang)

		f := Latitute()
		if f == 0 {
			t.Errorf("Latitude failed with lang %s", lang)
		}

		i := LatitudeDegreess()
		if i < -180 || i > 180 {
			t.Errorf("LatitudeDegress failed with lang %s", lang)
		}

		i = LatitudeMinutes()
		if i < 0 || i >= 60 {
			t.Errorf("LatitudeMinutes failed with lang %s", lang)
		}

		i = LatitudeSeconds()
		if i < 0 || i >= 60 {
			t.Errorf("LatitudeSeconds failed with lang %s", lang)
		}

		s := LatitudeDirection()
		if s != "N" && s != "S" {
			t.Errorf("LatitudeDirection failed with lang %s", lang)
		}

		f = Longitude()
		if f == 0 {
			t.Errorf("Longitude failed with lang %s", lang)
		}

		i = LongitudeDegrees()
		if i < -180 || i > 180 {
			t.Errorf("LongitudeDegrees failed with lang %s", lang)
		}

		i = LongitudeMinutes()
		if i < 0 || i >= 60 {
			t.Errorf("LongitudeMinutes failed with lang %s", lang)
		}

		i = LongitudeSeconds()
		if i < 0 || i >= 60 {
			t.Errorf("LongitudeSeconds failed with lang %s", lang)
		}

		s = LongitudeDirection()
		if s != "W" && s != "E" {
			t.Errorf("LongitudeDirection failed with lang %s", lang)
		}
	}
}
