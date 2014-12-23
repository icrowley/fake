package test

import (
	"testing"

	"github.com/icrowley/fake"
)

func TestInternet(t *testing.T) {
	fake.SetLang("en")

	v := fake.UserName()
	if v == "" {
		t.Error("UserName failed")
	}

	v = fake.TopLevelDomain()
	if v == "" {
		t.Error("TopLevelDomain failed")
	}

	v = fake.DomainName()
	if v == "" {
		t.Error("DomainName failed")
	}

	v = fake.EmailAddress()
	if v == "" {
		t.Error("EmailAddress failed")
	}

	v = fake.EmailSubject()
	if v == "" {
		t.Error("EmailSubject failed")
	}

	v = fake.EmailBody()
	if v == "" {
		t.Error("EmailBody failed")
	}

	v = fake.DomainZone()
	if v == "" {
		t.Error("DomainZone failed")
	}

	v = fake.IPv4()
	if v == "" {
		t.Error("IPv4 failed")
	}
}
