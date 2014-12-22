package fake

import "strconv"

func Continent() string {
	return lookup(lang, "continents", true)
}

func Country() string {
	return lookup(lang, "countries", true)
}

func City() string {
	city := lookup(lang, "cities", true)
	switch r.Intn(5) {
	case 0:
		return join(cityPrefix(), city)
	case 1:
		return join(city, citySuffix())
	default:
		return city
	}
}

func cityPrefix() string {
	return lookup(lang, "city_prefixes", false)
}

func citySuffix() string {
	return lookup(lang, "city_suffixes", false)
}

func State() string {
	return lookup(lang, "states", false)
}

func StateAbbrev() string {
	return lookup(lang, "state_abbrevs", false)
}

func Street() string {
	street := lookup(lang, "streets", true)
	return join(street, streetSuffix())
}

func StreetAddress() string {
	return join(Street(), strconv.Itoa(r.Intn(100)))
}

func streetSuffix() string {
	return lookup(lang, "street_suffixes", true)
}

func Zip() string {
	return generate(lang, "zip", true)
}

func Phone() string {
	return generate(lang, "phone", true)
}
