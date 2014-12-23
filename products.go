package fake

func Brand() string {
	if r.Intn(2) == 0 {
		return lookup(lang, "brands", true) + " " + lookup(lang, "brands", true)
	}
	return lookup(lang, "adjectives", true) + " " + lookup(lang, "brands", true)
}

func ProductName() string {
	productName := lookup(lang, "adjectives", true) + " " + lookup(lang, "nouns", true)
	if r.Intn(2) == 1 {
		productName = lookup(lang, "adjectives", true) + " " + productName
	}
	return productName
}

func Product() string {
	return Brand() + " " + ProductName()
}

func Model() string {
	seps := []string{"", " ", "-"}
	return CharactersN(r.Intn(3)+1) + seps[r.Intn(len(seps))] + Digits()
}
