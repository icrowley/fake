package fake

// Brand generates brand name
func Brand() string {
	return f.Brand()
}

// ProductName generates product name
func ProductName() string {
	return f.ProductName()
}

// Product generates product title as brand + product name
func Product() string {
	return f.Product()
}

// Model generates model name that consists of letters and digits, optionally with a hyphen between them
func Model() string {
	return f.Model()
}

// Brand generates brand name
func (f *Faker) Brand() string {
	return f.Company()
}

// ProductName generates product name
func (f *Faker) ProductName() string {
	productName := f.lookup(f.lang, "adjectives", true) + " " + f.lookup(f.lang, "nouns", true)
	if f.r.Intn(2) == 1 {
		productName = f.lookup(f.lang, "adjectives", true) + " " + productName
	}
	return productName
}

// Product generates product title as brand + product name
func (f *Faker) Product() string {
	return f.Brand() + " " + f.ProductName()
}

// Model generates model name that consists of letters and digits, optionally with a hyphen between them
func (f *Faker) Model() string {
	seps := []string{"", " ", "-"}
	return f.CharactersN(f.r.Intn(3)+1) + seps[f.r.Intn(len(seps))] + f.Digits()
}
