package faker_test

import (
	"github.com/icrowley/faker"

	"testing"
)

func TestFakerEn(t *testing.T) {
	f := faker.NewFaker(&faker.Config{Lang: "en"})
	name, err := f.FullName()
	if err != nil {
		t.Error(err)
	}
	if name == "" {
		t.Error("Name should not be blank")
	}
}

func TestFakerRuWithoutCallback(t *testing.T) {
	f := faker.NewFaker(&faker.Config{Lang: "ru"})
	name, err := f.FullName()
	if err == nil {
		t.Error("Faker call with no samples file should produce an error")
	}
	if name != "" {
		t.Error("Faker call with no samples should return blank string")
	}
}

func TestFakerRuWithCallback(t *testing.T) {
	f := faker.NewFaker(&faker.Config{Lang: "ru", EnFallback: true})
	name, err := f.FullName()
	if err != nil {
		t.Error("Faker call with no samples file with callback should not produce an error")
	}
	if name == "" {
		t.Error("Faker call with no samples with callback should not return blank string")
	}
}
