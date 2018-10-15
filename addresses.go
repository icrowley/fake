package fake

import "strconv"

// Continent generates random continent
func Continent() string {
	return f.Continent()
}

// Country generates random country
func Country() string {
	return f.Country()
}

// City generates random city
func City() string {
	return f.City()
}

// State generates random state
func State() string {
	return f.State()
}

// StateAbbrev generates random state abbreviation
func StateAbbrev() string {
	return f.StateAbbrev()
}

// Street generates random street name
func Street() string {
	return f.Street()
}

// StreetAddress generates random street name along with building number
func StreetAddress() string {
	return f.StreetAddress()
}

// Zip generates random zip code using one of the formats specifies in zip_format file
func Zip() string {
	return f.Zip()
}

// Phone generates random phone number using one of the formats format specified in phone_format file
func Phone() string {
	return f.Phone()
}

// Continent generates random continent
func (f *Faker) Continent() string {
	return f.lookup(f.lang, "continents", true)
}

// Country generates random country
func (f *Faker) Country() string {
	return f.lookup(f.lang, "countries", true)
}

// City generates random city
func (f *Faker) City() string {
	city := f.lookup(f.lang, "cities", true)
	switch f.r.Intn(5) {
	case 0:
		return join(f.cityPrefix(), city)
	case 1:
		return join(city, f.citySuffix())
	default:
		return city
	}
}

func (f *Faker) cityPrefix() string {
	return f.lookup(f.lang, "city_prefixes", false)
}

func (f *Faker) citySuffix() string {
	return f.lookup(f.lang, "city_suffixes", false)
}

// State generates random state
func (f *Faker) State() string {
	return f.lookup(f.lang, "states", false)
}

// StateAbbrev generates random state abbreviation
func (f *Faker) StateAbbrev() string {
	return f.lookup(f.lang, "state_abbrevs", false)
}

// Street generates random street name
func (f *Faker) Street() string {
	street := f.lookup(f.lang, "streets", true)
	return join(street, f.streetSuffix())
}

// StreetAddress generates random street name along with building number
func (f *Faker) StreetAddress() string {
	return join(f.Street(), strconv.Itoa(f.r.Intn(100)))
}

func (f *Faker) streetSuffix() string {
	return f.lookup(f.lang, "street_suffixes", true)
}

// Zip generates random zip code using one of the formats specifies in zip_format file
func (f *Faker) Zip() string {
	return f.generate(f.lang, "zips", true)
}

// Phone generates random phone number using one of the formats format specified in phone_format file
func (f *Faker) Phone() string {
	return f.generate(f.lang, "phones", true)
}
