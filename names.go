package fake

// MaleFirstName generates male first name
func MaleFirstName() string {
	return f.MaleFirstName()
}

// FemaleFirstName generates female first name
func FemaleFirstName() string {
	return f.FemaleFirstName()
}

// FirstName generates first name
func FirstName() string {
	return f.FirstName()
}

// MaleLastName generates male last name
func MaleLastName() string {
	return f.MaleLastName()
}

// FemaleLastName generates female last name
func FemaleLastName() string {
	return f.FemaleLastName()
}

// LastName generates last name
func LastName() string {
	return f.LastName()
}

// MalePatronymic generates male patronymic
func MalePatronymic() string {
	return f.MalePatronymic()
}

// FemalePatronymic generates female patronymic
func FemalePatronymic() string {
	return f.FemalePatronymic()
}

// Patronymic generates patronymic
func Patronymic() string {
	return f.Patronymic()
}

// MaleFullNameWithPrefix generates prefixed male full name
// if prefixes for the given language are available
func MaleFullNameWithPrefix() string {
	return f.MaleFullNameWithPrefix()
}

// FemaleFullNameWithPrefix generates prefixed female full name
// if prefixes for the given language are available
func FemaleFullNameWithPrefix() string {
	return f.FemaleFullNameWithPrefix()
}

// FullNameWithPrefix generates prefixed full name
// if prefixes for the given language are available
func FullNameWithPrefix() string {
	return f.FullNameWithPrefix()
}

// MaleFullNameWithSuffix generates suffixed male full name
// if suffixes for the given language are available
func MaleFullNameWithSuffix() string {
	return f.MaleFullNameWithSuffix()
}

// FemaleFullNameWithSuffix generates suffixed female full name
// if suffixes for the given language are available
func FemaleFullNameWithSuffix() string {
	return f.FemaleFullNameWithSuffix()
}

// FullNameWithSuffix generates suffixed full name
// if suffixes for the given language are available
func FullNameWithSuffix() string {
	return f.MaleFullNameWithSuffix()
}

// MaleFullName generates male full name
// it can occasionally include prefix or suffix
func MaleFullName() string {
	return f.MaleFullName()
}

// FemaleFullName generates female full name
// it can occasionally include prefix or suffix
func FemaleFullName() string {
	return f.FemaleFullName()
}

// FullName generates full name
// it can occasionally include prefix or suffix
func FullName() string {
	return f.FullName()
}

func (f *Faker) randGender() string {
	g := "male"
	if f.r.Intn(2) == 0 {
		g = "female"
	}
	return g
}

func (f *Faker) firstName(gender string) string {
	return f.lookup(f.lang, gender+"_first_names", true)
}

// MaleFirstName generates male first name
func (f *Faker) MaleFirstName() string {
	return f.firstName("male")
}

// FemaleFirstName generates female first name
func (f *Faker) FemaleFirstName() string {
	return f.firstName("female")
}

// FirstName generates first name
func (f *Faker) FirstName() string {
	return f.firstName(f.randGender())
}

func (f *Faker) lastName(gender string) string {
	return f.lookup(f.lang, gender+"_last_names", true)
}

// MaleLastName generates male last name
func (f *Faker) MaleLastName() string {
	return f.lastName("male")
}

// FemaleLastName generates female last name
func (f *Faker) FemaleLastName() string {
	return f.lastName("female")
}

// LastName generates last name
func (f *Faker) LastName() string {
	return f.lastName(f.randGender())
}

func (f *Faker) patronymic(gender string) string {
	return f.lookup(f.lang, gender+"_patronymics", false)
}

// MalePatronymic generates male f.patronymic
func (f *Faker) MalePatronymic() string {
	return f.patronymic("male")
}

// FemalePatronymic generates female f.patronymic
func (f *Faker) FemalePatronymic() string {
	return f.patronymic("female")
}

// Patronymic generates f.patronymic
func (f *Faker) Patronymic() string {
	return f.patronymic(f.randGender())
}

func (f *Faker) prefix(gender string) string {
	return f.lookup(f.lang, gender+"_name_prefixes", false)
}

func (f *Faker) suffix(gender string) string {
	return f.lookup(f.lang, gender+"_name_suffixes", false)
}

func (f *Faker) fullNameWithPrefix(gender string) string {
	return join(f.prefix(gender), f.firstName(gender), f.lastName(gender))
}

// MaleFullNameWithPrefix generates prefixed male full name
// if prefixes for the given f.language are available
func (f *Faker) MaleFullNameWithPrefix() string {
	return f.fullNameWithPrefix("male")
}

// FemaleFullNameWithPrefix generates prefixed female full name
// if prefixes for the given f.language are available
func (f *Faker) FemaleFullNameWithPrefix() string {
	return f.fullNameWithPrefix("female")
}

// FullNameWithPrefix generates prefixed full name
// if prefixes for the given f.language are available
func (f *Faker) FullNameWithPrefix() string {
	return f.fullNameWithPrefix(f.randGender())
}

func (f *Faker) fullNameWithSuffix(gender string) string {
	return join(f.firstName(gender), f.lastName(gender), f.suffix(gender))
}

// MaleFullNameWithSuffix generates suffixed male full name
// if suffixes for the given f.language are available
func (f *Faker) MaleFullNameWithSuffix() string {
	return f.fullNameWithPrefix("male")
}

// FemaleFullNameWithSuffix generates suffixed female full name
// if suffixes for the given f.language are available
func (f *Faker) FemaleFullNameWithSuffix() string {
	return f.fullNameWithPrefix("female")
}

// FullNameWithSuffix generates suffixed full name
// if suffixes for the given f.language are available
func (f *Faker) FullNameWithSuffix() string {
	return f.fullNameWithPrefix(f.randGender())
}

func (f *Faker) fullName(gender string) string {
	switch f.r.Intn(10) {
	case 0:
		return f.fullNameWithPrefix(gender)
	case 1:
		return f.fullNameWithSuffix(gender)
	default:
		return join(f.firstName(gender), f.lastName(gender))
	}
}

// MaleFullName generates male full name
// it can occasionally include prefix or suffix
func (f *Faker) MaleFullName() string {
	return f.fullName("male")
}

// FemaleFullName generates female full name
// it can occasionally include prefix or suffix
func (f *Faker) FemaleFullName() string {
	return f.fullName("female")
}

// FullName generates full name
// it can occasionally include prefix or suffix
func (f *Faker) FullName() string {
	return f.fullName(f.randGender())
}
