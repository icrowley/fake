package fake_test

import (
	"testing"

	"github.com/icrowley/fake"
)

func TestNames(t *testing.T) {
	fake.SetLang("en")

	v := fake.MaleFirstName()
	if v == "" {
		t.Error("MaleFirstName failed")
	}

	v = fake.FemaleFirstName()
	if v == "" {
		t.Error("FemaleFirstName failed")
	}

	v = fake.FirstName()
	if v == "" {
		t.Error("FirstName failed")
	}

	v = fake.MaleLastName()
	if v == "" {
		t.Error("MaleLastName failed")
	}

	v = fake.FemaleLastName()
	if v == "" {
		t.Error("FemaleLastName failed")
	}

	v = fake.LastName()
	if v == "" {
		t.Error("LastName failed")
	}

	v = fake.MalePatronymic()
	if v == "" {
		t.Error("MalePatronymic failed")
	}

	v = fake.FemalePatronymic()
	if v == "" {
		t.Error("FemalePatronymic failed")
	}

	v = fake.Patronymic()
	if v == "" {
		t.Error("Patronymic failed")
	}

	v = fake.MaleFullNameWithPrefix()
	if v == "" {
		t.Error("MaleFullNameWithPrefix failed")
	}

	v = fake.FemaleFullNameWithPrefix()
	if v == "" {
		t.Error("FemaleFullNameWithPrefix failed")
	}

	v = fake.FullNameWithPrefix()
	if v == "" {
		t.Error("FullNameWithPrefix failed")
	}

	v = fake.MaleFullNameWithSuffix()
	if v == "" {
		t.Error("MaleFullNameWithSuffix failed")
	}

	v = fake.FemaleFullNameWithSuffix()
	if v == "" {
		t.Error("FemaleFullNameWithSuffix failed")
	}

	v = fake.FullNameWithSuffix()
	if v == "" {
		t.Error("FullNameWithSuffix failed")
	}

	v = fake.MaleFullName()
	if v == "" {
		t.Error("MaleFullName failed")
	}

	v = fake.FemaleFullName()
	if v == "" {
		t.Error("FemaleFullName failed")
	}

	v = fake.FullName()
	if v == "" {
		t.Error("FullName failed")
	}
}
