package fake_test

import (
	"github.com/icrowley/fake"

	"testing"
)

func TestCheckLang(t *testing.T) {
	_, err := fake.NewFaker(fake.Config{Lang: "ru"})
	if err == nil {
		t.Error("NewFaker with nonexistent lang should return error")
	}

	_, err = fake.NewFaker(fake.Config{})
	if err != nil {
		t.Error("NewFaker without lang should use en")
	}
}

func TestFakerEn(t *testing.T) {
	f, err := fake.NewFaker(fake.Config{Lang: "en"})
	name, err := f.FullName()
	if err != nil {
		t.Error(err)
	}
	if name == "" {
		t.Error("Name should not be blank")
	}
}

func TestFakerCaWithoutCallback(t *testing.T) {
	f, _ := fake.NewFaker(fake.Config{Lang: "ca"})
	name, err := f.FullName()
	if err == nil {
		t.Error("Faker call with no samples file should produce an error")
	}
	if name != "" {
		t.Error("Faker call with no samples should return blank string")
	}
}

func TestFakerCaWithCallback(t *testing.T) {
	f, _ := fake.NewFaker(fake.Config{Lang: "ca", EnFallback: true})
	name, err := f.FullName()
	if err != nil {
		t.Error("Faker call with no samples file with callback should not produce an error")
	}
	if name == "" {
		t.Error("Faker call with no samples with callback should not return blank string")
	}
}

func TestDefaults(t *testing.T) {
	name := fake.FullName()
	if name == "" {
		t.Error("Default faker call should return a string")
	}
}

func TestDefaultCa(t *testing.T) {
	fake.SetLang("ca")
	name := fake.FullName()
	if name != "" {
		t.Errorf("Default faker call should return blank string with no existing samples, but returned %s\n", name)
	}
}
