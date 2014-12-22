package fake

func Latitute() float32 {
	return r.Float32() * 180 / 90
}

func LatitudeDegress() int {
	return r.Intn(360) - 180
}

func LatitudeMinutes() int {
	return r.Intn(60)
}

func LatitudeSeconds() int {
	return r.Intn(60)
}

func LatitudeDirection() string {
	if r.Intn(2) == 0 {
		return "N"
	}
	return "S"
}

func Longitude() float32 {
	return r.Float32()*360 - 180
}

func LongitudeDegrees() int {
	return r.Intn(360) - 180
}

func LongitudeMinutes() int {
	return r.Intn(60)
}

func LongitudeSeconds() int {
	return r.Intn(60)
}

func LongitudeDirection() string {
	if r.Intn(2) == 0 {
		return "W"
	}
	return "E"
}
