package fake_test

import (
	"github.com/icrowley/fake"

	"testing"
)

func TestSetLang(t *testing.T) {
	err := fake.SetLang("it")
	if err != nil {
		t.Error("SetLang should successfully set lang")
	}

	err = fake.SetLang("sd")
	if err == nil {
		t.Error("SetLang with nonexistent lang should return error")
	}
}

func TestFakerEn(t *testing.T) {
	fake.SetLang("en")
	name, err := fake.FullName()
	if err != nil {
		t.Error(err.Error())
	}
	if name == "" {
		t.Error("Name should not be blank")
	}
}

func TestFakerCaWithoutCallback(t *testing.T) {
	fake.SetLang("ca")
	name, err := fake.FullName()
	if err == nil {
		t.Error("Fake call with no samples file should produce an error")
	}
	if name != "" {
		t.Error("Fake call with no samples should return blank string")
	}
}

func TestFakerCaWithCallback(t *testing.T) {
	fake.SetLang("ca")
	fake.EnFallback(true)
	name, err := fake.FullName()
	if err != nil {
		t.Error("Fake call with no samples file with callback should not produce an error")
	}
	if name == "" {
		t.Error("Fake call with no samples with callback should not return blank string")
	}
}
