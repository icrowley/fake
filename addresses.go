package fake

import "strconv"

func Continent() string {
	return lookup("addresses", "continents", lang, true)
}

func Country() string {
	return lookup("addresses", "countries", lang, true)
}

func City() string {
	city := lookup("addresses", "cities", lang, true)
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
	return lookup("addresses", "city_prefixes", lang, false)
}

func citySuffix() string {
	return lookup("addresses", "city_suffixes", lang, false)
}

func State() string {
	return lookup("addresses", "states", lang, false)
}

func StateAbbrev() string {
	return lookup("addresses", "state_abbrevs", lang, false)
}

func Street() string {
	street := lookup("addresses", "streets", lang, true)
	return join(street, streetSuffix())
}

func StreetAddress() string {
	return join(Street(), strconv.Itoa(r.Intn(100)))
}

func streetSuffix() string {
	return lookup("addresses", "street_suffixes", lang, true)
}

func Zip() string {
	return generate("addresses", "zip", lang, true)
}

func Phone() string {
	return generate("addresses", "phone", lang, true)
}
