package fake

func randGender() string {
	g := "male"
	if r.Intn(2) == 0 {
		g = "female"
	}
	return g
}

func firstName(gender string) string {
	return lookup("name", gender+"_first_names", lang, true)
}

func MaleFirstName() string {
	return firstName("male")
}

func FemaleFirstName() string {
	return firstName("female")
}

func FirstName() string {
	return firstName(randGender())
}

func lastName(gender string) string {
	return lookup("name", gender+"_last_names", lang, true)
}

func MaleLastName() string {
	return lastName("male")
}

func FemaleLastName() string {
	return lastName("female")
}

func LastName() string {
	return lastName(randGender())
}

func patronymic(gender string) string {
	return lookup("name", gender+"_patronymics", lang, false)
}

func MalePatronymic() string {
	return patronymic("male")
}

func FemalePatronymic() string {
	return patronymic("female")
}

func Patronymic() string {
	return patronymic(randGender())
}

func prefix(gender string) string {
	return lookup("name", gender+"_prefixes", lang, false)
}

func suffix(gender string) string {
	return lookup("name", gender+"_suffixes", lang, false)
}

func fullNameWithPrefix(gender string) string {
	return join(prefix(gender), firstName(gender), lastName(gender))
}

func MaleFullNameWithPrefix() string {
	return fullNameWithPrefix("male")
}

func FemaleFullNameWithPrefix() string {
	return fullNameWithPrefix("female")
}

func fullNameWithSuffix(gender string) string {
	return join(firstName(gender), lastName(gender), suffix(gender))
}

func MaleFullNameWithSuffix() string {
	return fullNameWithPrefix("male")
}

func FemaleFullNameWithSuffix() string {
	return fullNameWithPrefix("female")
}

func FullNameWithSuffix() string {
	return fullNameWithPrefix(randGender())
}

func fullName(gender string) string {
	switch r.Intn(10) {
	case 0:
		return fullNameWithPrefix(gender)
	case 1:
		return fullNameWithSuffix(gender)
	default:
		return join(firstName(gender), lastName(gender))
	}
}

func MaleFullName() string {
	return fullName("male")
}

func FemaleFullName() string {
	return fullName("female")
}

func FullName() string {
	return fullName(randGender())
}
