package fake

func Company() string {
	return lookup("job", "companies", lang, true)
}

func JobTitle() string {
	job := lookup("job", "jobs", lang, true)
	return join(job, jobTitleSuffix())
}

func jobTitleSuffix() string {
	return lookup("job", "jobs", lang, false)
}

func Industry() string {
	return lookup("job", "industries", lang, true)
}
