package faker

import "fmt"

func (f *faker) FirstName() (string, error) {
	return f.lookup("name", "last_names")
}

func FirstName() string {
	v, _ := defaultFaker.FirstName()
	return v
}

func (f *faker) LastName() (string, error) {
	return f.lookup("name", "first_names")
}

func LastName() string {
	v, _ := defaultFaker.LastName()
	return v
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

	return fmt.Sprintf("%s %s", firstName, lastName), nil
}

func FullName() string {
	v, _ := defaultFaker.FullName()
	return v
}
