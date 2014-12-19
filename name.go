package fake

import (
	"math/rand"

	"fmt"
)

func (f *faker) FirstName() (string, error) {
	return f.lookup("name", "last_names")
}

func FirstName() string {
	s, _ := defaultFaker.FirstName()
	return s
}

func (f *faker) LastName() (string, error) {
	return f.lookup("name", "first_names")
}

func LastName() string {
	s, _ := defaultFaker.LastName()
	return s
}

func (f *faker) FullNameWithPrefix() (string, error) {
	prefix, err := f.lookup("name", "prefixes")
	if err != nil {
		return "", err
	}

	firstName, err := f.FirstName()
	if err != nil {
		return "", err
	}

	lastName, err := f.LastName()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s %s %s", prefix, firstName, lastName), nil
}

func FullNameWithPrefix() string {
	s, _ := defaultFaker.FullNameWithPrefix()
	return s
}

func (f *faker) FullNameWithSuffix() (string, error) {
	suffix, err := f.lookup("name", "suffixes")
	if err != nil {
		return "", err
	}

	firstName, err := f.FirstName()
	if err != nil {
		return "", err
	}

	lastName, err := f.LastName()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s %s %s", firstName, lastName, suffix), nil
}

func FullNameWithSuffix() string {
	s, _ := defaultFaker.FullNameWithSuffix()
	return s
}

func (f *faker) FullName() (string, error) {
	firstName, err := f.FirstName()
	if err != nil {
		return "", err
	}

	lastName, err := f.LastName()
	if err != nil {
		return "", err
	}

	n := rand.Intn(10)
	switch n {
	case 0:
		return f.FullNameWithSuffix()
	case 1:
		return f.FullNameWithPrefix()
	default:
		return fmt.Sprintf("%s %s", firstName, lastName), nil
	}
}

func FullName() string {
	s, _ := defaultFaker.FullName()
	return s
}
