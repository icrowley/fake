package fake

import (
	"time"
)

func Day() int {
	return r.Intn(31) + 1
}

func WeekDay() string {
	return lookup(lang, "weekdays", true)
}

func WeekDayShort() string {
	return lookup(lang, "weekdays_short", true)
}

func WeekdayNum() int {
	return r.Intn(7) + 1
}

func Month() string {
	return lookup(lang, "months", true)
}

func MonthShort() string {
	return lookup(lang, "months_short", true)
}

func MonthNum() int {
	return r.Intn(12) + 1
}

func Year(fromOffset, toOffset int) int {
	n := r.Intn(fromOffset+toOffset) + 1
	return (time.Now().Year() - fromOffset) + n
}
