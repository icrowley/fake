package fake_test

import (
	"testing"

	"github.com/icrowley/fake"
)

func TestJobs(t *testing.T) {
	fake.SetLang("en")

	v := fake.Company()
	if v == "" {
		t.Error("Company failed")
	}

	v = fake.JobTitle()
	if v == "" {
		t.Error("JobTitle failed")
	}

	v = fake.Industry()
	if v == "" {
		t.Error("Industry failed")
	}
}
