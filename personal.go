package fake

func Gender() string {
	return lookup("personal", "genders", lang, true)
}

func GenderAbbrev() string {
	g := Gender()
	if g != "" {
		return string(g[0])
	}
	return ""
}

func Language() string {
	return lookup("personal", "languages", lang, true)
}
