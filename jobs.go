package fake

func Company() string {
	return lookup(lang, "companies", true)
}

func JobTitle() string {
	job := lookup(lang, "jobs", true)
	return join(job, jobTitleSuffix())
}

func jobTitleSuffix() string {
	return lookup(lang, "jobs_suffixes", false)
}

func Industry() string {
	return lookup(lang, "industries", true)
}
