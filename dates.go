package fake

// Day generates day of the month
func Day() int {
	return f.Day()
}

// WeekDay generates name ot the week day
func WeekDay() string {
	return f.WeekDay()
}

// WeekDayShort generates abbreviated name of the week day
func WeekDayShort() string {
	return f.WeekDayShort()
}

// WeekdayNum generates number of the day of the week
func WeekdayNum() int {
	return f.WeekdayNum()
}

// Month generates month name
func Month() string {
	return f.Month()
}

// MonthShort generates abbreviated month name
func MonthShort() string {
	return f.MonthShort()
}

// MonthNum generates month number (from 1 to 12)
func MonthNum() int {
	return f.MonthNum()
}

// Year generates year using the given boundaries
func Year(from, to int) int {
	return f.Year(from, to)
}

// Day generates day of the month
func (f *Faker) Day() int {
	return f.r.Intn(31) + 1
}

// WeekDay generates name ot the week day
func (f *Faker) WeekDay() string {
	return f.lookup(f.lang, "weekdays", true)
}

// WeekDayShort generates abbreviated name of the week day
func (f *Faker) WeekDayShort() string {
	return f.lookup(f.lang, "weekdays_short", true)
}

// WeekdayNum generates number of the day of the week
func (f *Faker) WeekdayNum() int {
	return f.r.Intn(7) + 1
}

// Month generates month name
func (f *Faker) Month() string {
	return f.lookup(f.lang, "months", true)
}

// MonthShort generates abbreviated month name
func (f *Faker) MonthShort() string {
	return f.lookup(f.lang, "months_short", true)
}

// MonthNum generates month number (from 1 to 12)
func (f *Faker) MonthNum() int {
	return f.r.Intn(12) + 1
}

// Year generates year using the given boundaries
func (f *Faker) Year(from, to int) int {
	n := f.r.Intn(to-from) + 1
	return from + n
}
