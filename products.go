package fake

func Brand() string {
	return Company()
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
