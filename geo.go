package fake

// Latitude generates latitude (from -90.0 to 90.0)
func Latitude() float32 {
	return f.Latitude()
}

// LatitudeDegrees generates latitude degrees (from -90 to 90)
func LatitudeDegrees() int {
	return f.LatitudeDegrees()
}

// LatitudeMinutes generates latitude minutes (from 0 to 60)
func LatitudeMinutes() int {
	return f.LatitudeMinutes()
}

// LatitudeSeconds generates latitude seconds (from 0 to 60)
func LatitudeSeconds() int {
	return f.LatitudeSeconds()
}

// LatitudeDirection generates latitude direction (N(orth) o S(outh))
func LatitudeDirection() string {
	return f.LatitudeDirection()
}

// Longitude generates longitude (from -180 to 180)
func Longitude() float32 {
	return f.Longitude()
}

// LongitudeDegrees generates longitude degrees (from -180 to 180)
func LongitudeDegrees() int {
	return f.LongitudeDegrees()
}

// LongitudeMinutes generates (from 0 to 60)
func LongitudeMinutes() int {
	return f.LongitudeMinutes()
}

// LongitudeSeconds generates (from 0 to 60)
func LongitudeSeconds() int {
	return f.LongitudeSeconds()
}

// LongitudeDirection generates (W(est) or E(ast))
func LongitudeDirection() string {
	return f.LongitudeDirection()
}

// Latitude generates latitude (from -90.0 to 90.0)
func (f *Faker) Latitude() float32 {
	return f.r.Float32()*180 - 90
}

// LatitudeDegrees generates latitude degrees (from -90 to 90)
func (f *Faker) LatitudeDegrees() int {
	return f.r.Intn(180) - 90
}

// LatitudeMinutes generates latitude minutes (from 0 to 60)
func (f *Faker) LatitudeMinutes() int {
	return f.r.Intn(60)
}

// LatitudeSeconds generates latitude seconds (from 0 to 60)
func (f *Faker) LatitudeSeconds() int {
	return f.r.Intn(60)
}

// LatitudeDirection generates latitude direction (N(orth) o S(outh))
func (f *Faker) LatitudeDirection() string {
	if f.r.Intn(2) == 0 {
		return "N"
	}
	return "S"
}

// Longitude generates longitude (from -180 to 180)
func (f *Faker) Longitude() float32 {
	return f.r.Float32()*360 - 180
}

// LongitudeDegrees generates longitude degrees (from -180 to 180)
func (f *Faker) LongitudeDegrees() int {
	return f.r.Intn(360) - 180
}

// LongitudeMinutes generates (from 0 to 60)
func (f *Faker) LongitudeMinutes() int {
	return f.r.Intn(60)
}

// LongitudeSeconds generates (from 0 to 60)
func (f *Faker) LongitudeSeconds() int {
	return f.r.Intn(60)
}

// LongitudeDirection generates (W(est) or E(ast))
func (f *Faker) LongitudeDirection() string {
	if f.r.Intn(2) == 0 {
		return "W"
	}
	return "E"
}
