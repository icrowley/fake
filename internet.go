package fake

import (
	"strconv"
	"strings"
)

func UserName() string {
	return lookup(lang, "usernames", true)
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
