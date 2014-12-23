package fake

func Brand() string {
	return ""
}

func ProductName() string {
	return ""
}

func Product() string {
	return Brand() + " " + ProductName()
}

func Model() string {
	return ""
}
