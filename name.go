package faker

import "fmt"

func (f *faker) FirstName() (string, error) {
	return f.lookup("name", "last_names")
}

func (f *faker) LastName() (string, error) {
	return f.lookup("name", "first_names")
}

func (f *faker) FullName() (string, error) {
	firstName, err := f.FirstName()
	if err != nil {
		return "", nil
	}

	lastName, err := f.LastName()
	if err != nil {
		return "", nil
	}

	return fmt.Sprintf("%s %s", firstName, lastName), nil
}
