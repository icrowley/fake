package fake_test

import (
	"testing"
	"time"

	"github.com/icrowley/fake"
)

func TestDates(t *testing.T) {
	fake.SetLang("en")

	v := fake.WeekDay()
	if v == "" {
		t.Error("WeekDay failed")
	}

	v = fake.WeekDayShort()
	if v == "" {
		t.Error("WeekDayShort failed")
	}

	n := fake.WeekdayNum()
	if n < 0 || n > 7 {
		t.Error("WeekdayNum failed")
	}

	v = fake.Month()
	if v == "" {
		t.Error("Month failed")
	}

	v = fake.MonthShort()
	if v == "" {
		t.Error("MonthShort failed")
	}

	n = fake.MonthNum()
	if n < 0 || n > 31 {
		t.Error("MonthNum failed")
	}

	n = fake.Year(50, 20)
	year := time.Now().Year()
	if n < year-50 || n > year+20 {
		t.Error("Year failed")
	}
}
