package fake

import (
	"net"
	"strings"

	"github.com/corpix/uarand"
)

// UserName generates user name in one of the following forms
// first name + last name, letter + last names or concatenation of from 1 to 3 lowercased words
func UserName() string {
	return f.UserName()
}

// TopLevelDomain generates random top level domain
func TopLevelDomain() string {
	return f.TopLevelDomain()
}

// DomainName generates random domain name
func DomainName() string {
	return f.DomainName()
}

// EmailAddress generates email address
func EmailAddress() string {
	return f.EmailAddress()
}

// EmailSubject generates random email subject
func EmailSubject() string {
	return f.EmailSubject()
}

// EmailBody generates random email body
func EmailBody() string {
	return f.EmailBody()
}

// DomainZone generates random domain zone
func DomainZone() string {
	return f.DomainZone()
}

// IPv4 generates IPv4 address
func IPv4() string {
	return f.IPv4()
}

// IPv6 generates IPv6 address
func IPv6() string {
	return f.IPv6()
}

// UserAgent generates a random user agent.
func UserAgent() string {
	return f.UserAgent()
}

// UserName generates user name in one of the following forms
// first name + last name, letter + last names or concatenation of from 1 to 3 lowercased words
func (f *Faker) UserName() string {
	gender := f.randGender()
	switch f.r.Intn(3) {
	case 0:
		return f.lookup("en", gender+"_first_names", false) + f.lookup(f.lang, gender+"_last_names", false)
	case 1:
		return Character() + f.lookup(f.lang, gender+"_last_names", false)
	default:
		return strings.Replace(WordsN(f.r.Intn(3)+1), " ", "_", -1)
	}
}

// TopLevelDomain generates random top level domain
func (f *Faker) TopLevelDomain() string {
	return f.lookup(f.lang, "top_level_domains", true)
}

// DomainName generates random domain name
func (f *Faker) DomainName() string {
	return Company() + "." + TopLevelDomain()
}

// EmailAddress generates email address
func (f *Faker) EmailAddress() string {
	return UserName() + "@" + DomainName()
}

// EmailSubject generates random email subject
func (f *Faker) EmailSubject() string {
	return Sentence()
}

// EmailBody generates random email body
func (f *Faker) EmailBody() string {
	return Paragraphs()
}

// DomainZone generates random domain zone
func (f *Faker) DomainZone() string {
	return f.lookup(f.lang, "domain_zones", true)
}

// IPv4 generates IPv4 address
func (f *Faker) IPv4() string {
	size := 4
	ip := make([]byte, size)
	for i := 0; i < size; i++ {
		ip[i] = byte(f.r.Intn(256))
	}
	return net.IP(ip).To4().String()
}

// IPv6 generates IPv6 address
func (f *Faker) IPv6() string {
	size := 16
	ip := make([]byte, size)
	for i := 0; i < size; i++ {
		ip[i] = byte(f.r.Intn(256))
	}
	return net.IP(ip).To16().String()
}

// UserAgent generates a random user agent.
func (f *Faker) UserAgent() string {
	return uarand.GetRandom()
}
