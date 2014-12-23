package test

import (
	"testing"
	"time"

	"github.com/icrowley/fake"
)

func TestDates(t *testing.T) {
	langs := []string{"en", "ru"}
	for _, lang := range langs {
		fake.SetLang(lang)

		v := fake.WeekDay()
		if v == "" {
			t.Errorf("WeekDay failed with lang %s", lang)
		}

		v = fake.WeekDayShort()
		if v == "" {
			t.Errorf("WeekDayShort failed with lang %s", lang)
		}

		n := fake.WeekdayNum()
		if n < 0 || n > 7 {
			t.Errorf("WeekdayNum failed with lang %s", lang)
		}

		v = fake.Month()
		if v == "" {
			t.Errorf("Month failed with lang %s", lang)
		}

		v = fake.MonthShort()
		if v == "" {
			t.Errorf("MonthShort failed with lang %s", lang)
		}

		n = fake.MonthNum()
		if n < 0 || n > 31 {
			t.Errorf("MonthNum failed with lang %s", lang)
		}

		n = fake.Year(50, 20)
		year := time.Now().Year()
		if n < year-50 || n > year+20 {
			t.Errorf("Year failed with lang %s", lang)
		}
	}
}
