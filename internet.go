package fake

import (
	"strconv"
	"strings"
)

func UserName() string {
	gender := randGender()
	switch r.Intn(2) {
	case 0:
		return lookup("en", gender+"_first_names", false) + lookup(lang, gender+"_last_names", false)
	case 1:
		return Character() + lookup(lang, gender+"_last_names", false)
	default:
		return strings.Replace(WordsN(r.Intn(3)+1), " ", "_", -1)
	}
}

func TopLevelDomain() string {
	return lookup(lang, "top_level_domains", true)
}

func DomainName() string {
	return Company() + "." + TopLevelDomain()
}

func EmailAddress() string {
	return UserName() + "@" + DomainName()
}

func EmailSubject() string {
	return Sentence()
}

func EmailBody() string {
	return Paragraphs()
}

func DomainZone() string {
	return lookup(lang, "domain_zones", true)
}

func IPv4() string {
	ip := make([]string, 4)
	for i := 0; i < 4; i++ {
		ip[i] = strconv.Itoa(r.Intn(256))
	}
	return strings.Join(ip, ".")
}
