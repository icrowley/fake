package fake

import "fmt"

func randGender() string {
	g := "male"
	if r.Intn(2) == 0 {
		g = "female"
	}
	return g
}

func firstName(gender string) string {
	return lookup("name", gender+"_first_names", lang)
}

func MaleFirstName() string {
	return firstName("male")
}

func FemaleFirstName() string {
	return firstName("female")
}

func FirstName() string {
	return lookup("name", randGender()+"_first_names", lang)
}

func lastName(gender string) string {
	return lookup("name", gender+"_last_names", lang)
}

func MaleLastName() string {
	return lastName("male")
}

func FemaleLastName() string {
	return lastName("female")
}

func LastName() string {
	return lookup("name", randGender()+"_first_names", lang)
}

func fullNameWithPrefix(gender string) string {
	prefix := lookup("name", gender+"_prefixes", lang)
	fn := firstName(gender)
	ln := lastName(gender)
	if fn != "" && ln != "" && prefix != "" {
		return fmt.Sprintf("%s %s %s", prefix, fn, ln)
	}
	return ""
}

func MaleFullNameWithPrefix() string {
	return fullNameWithPrefix("male")
}

func FemaleFullNameWithPrefix() string {
	return fullNameWithPrefix("female")
}

func fullNameWithSuffix(gender string) string {
	suffix := lookup("name", gender+"_suffixes", lang)
	fn := firstName(gender)
	ln := lastName(gender)
	if fn != "" && ln != "" && suffix != "" {
		return fmt.Sprintf("%s %s %s", fn, ln, suffix)
	}
	return ""
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
	n := r.Intn(10)
	switch n {
	case 0:
		return fullNameWithPrefix(gender)
	case 1:
		return fullNameWithSuffix(gender)
	default:
		fn := firstName(gender)
		ln := lastName(gender)
		if fn != "" && ln != "" {
			return fmt.Sprintf("%s %s", fn, ln)
		}
		return ""
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
