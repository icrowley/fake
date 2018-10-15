package fake

import (
	"strings"
)

// Company generates company name
func Company() string {
	return f.Company()
}

// JobTitle generates job title
func JobTitle() string {
	return f.JobTitle()
}

// Industry generates industry name
func Industry() string {
	return f.Industry()
}

// Company generates company name
func (f *Faker) Company() string {
	return f.lookup(f.lang, "companies", true)
}

// JobTitle generates job title
func (f *Faker) JobTitle() string {
	job := f.lookup(f.lang, "jobs", true)
	return strings.Replace(job, "#{N}", f.jobTitleSuffix(), 1)
}

func (f *Faker) jobTitleSuffix() string {
	return f.lookup(f.lang, "jobs_suffixes", false)
}

// Industry generates industry name
func (f *Faker) Industry() string {
	return f.lookup(f.lang, "industries", true)
}
