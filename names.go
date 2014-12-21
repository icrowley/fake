package fake

import "fmt"

func randGender() string {
	g := "male"
	if r.Intn(2) == 0 {
		g = "female"
	}
	return g
}

func firstName(gender string) (string, error) {
	return lookup("name", gender+"_first_names", lang)
}

func MaleFirstName() (string, error) {
	return firstName("male")
}

func FemaleFirstName() (string, error) {
	return firstName("female")
}

func FirstName() (string, error) {
	return lookup("name", randGender()+"_first_names", lang)
}

func lastName(gender string) (string, error) {
	return lookup("name", gender+"_last_names", lang)
}

func MaleLastName() (string, error) {
	return lastName("male")
}

func FemaleLastName() (string, error) {
	return lastName("female")
}

func LastName() (string, error) {
	return lookup("name", randGender()+"_first_names", lang)
}

func fullNameWithPrefix(gender string) (string, error) {
	prefix, err := lookup("name", gender+"_prefixes", lang)
	if err != nil {
		return "", err
	}

	fn, err := firstName(gender)
	if err != nil {
		return "", err
	}

	ln, err := lastName(gender)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s %s %s", prefix, fn, ln), nil
}

func MaleFullNameWithPrefix() (string, error) {
	return fullNameWithPrefix("male")
}

func FemaleFullNameWithPrefix() (string, error) {
	return fullNameWithPrefix("female")
}

func fullNameWithSuffix(gender string) (string, error) {
	suffix, err := lookup("name", gender+"_suffixes", lang)
	if err != nil {
		return "", err
	}

	fn, err := firstName(gender)
	if err != nil {
		return "", err
	}

	ln, err := lastName(gender)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s %s %s", fn, ln, suffix), nil
}

func MaleFullNameWithSuffix() (string, error) {
	return fullNameWithPrefix("male")
}

func FemaleFullNameWithSuffix() (string, error) {
	return fullNameWithPrefix("female")
}

func FullNameWithSuffix() (string, error) {
	return fullNameWithPrefix(randGender())
}

func fullName(gender string) (string, error) {
	n := r.Intn(10)
	switch n {
	case 0:
		return fullNameWithPrefix(gender)
	case 1:
		return fullNameWithSuffix(gender)
	default:
		fn, err := firstName(gender)
		if err != nil {
			return "", err
		}

		ln, err := lastName(gender)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("%s %s", fn, ln), nil
	}
}

func MaleFullName() (string, error) {
	return fullName("male")
}

func FemaleFullName() (string, error) {
	return fullName("female")
}

func FullName() (string, error) {
	return fullName(randGender())
}
