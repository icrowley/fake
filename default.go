package faker

var defaultFaker = &faker{lang: "en", samples: make(samplesMap)}

func SetLang(lang string) {
	defaultFaker.lang = lang
}

func UseLocalData(flag bool) {
	defaultFaker.useLocalData = flag
}

func enFallback(flag bool) {
	defaultFaker.enFallback = flag
}
