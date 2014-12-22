package fake

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// cat/subcat/lang/samples
type samplesTree map[string]map[string][]string

var samplesCache = make(samplesTree)
var r = rand.New(rand.NewSource(time.Now().UnixNano()))
var lang = "en"
var useLocalData = false
var enFallback = false

var (
	ErrInsufficientParams = fmt.Errorf("Insufficient params to lookup the samples")
	ErrNoLanguageFn       = func(lang string) error { return fmt.Errorf("The language passed (%s) is not available", lang) }
	ErrNoSamplesFn        = func(lang string) error { return fmt.Errorf("No samples found for language: %s", lang) }
)

func SetLang(newLang string) error {
	availLangs := []string{"en", "br", "ca", "da", "de", "fi", "fr", "mx", "nl", "se", "sn", "uk", "us", "ja", "ar", "cn", "cs",
		"ga", "it", "kr", "nb", "ph", "th", "vn"}

	found := false
	for _, l := range availLangs {
		if newLang == l {
			found = true
			break
		}
	}
	if !found {
		return ErrNoLanguageFn(newLang)
	}
	lang = newLang
	return nil
}

func UseLocalData(flag bool) {
	useLocalData = flag
}

func EnFallback(flag bool) {
	enFallback = flag
}

func (st samplesTree) hasKeyPath(lang, cat string) bool {
	if _, ok := st[lang]; ok {
		if _, ok = st[lang][cat]; ok {
			return true
		}
	}
	return false
}

func join(parts ...string) string {
	var filtered []string
	for _, part := range parts {
		if part != "" {
			filtered = append(filtered, part)
		}
	}
	return strings.Join(filtered, " ")
}

func generate(lag, cat string, fallback bool) string {
	format := lookup(lang, cat+"_format", fallback)
	var result string
	for _, ru := range format {
		if ru != '#' {
			result += string(ru)
		} else {
			result += strconv.Itoa(r.Intn(10))
		}
	}
	return result
}

func lookup(lang, cat string, fallback bool) string {
	var samples []string

	if samplesCache.hasKeyPath(lang, cat) {
		samples = samplesCache[lang][cat]
	} else {
		var err error
		samples, err = populateSamples(lang, cat)
		if err != nil {
			if lang != "en" && fallback && enFallback && err.Error() == ErrNoSamplesFn(lang).Error() {
				return lookup("en", cat, false)
			}
			return ""
		}
	}

	return samples[r.Intn(len(samples))]
}

func populateSamples(lang, cat string) ([]string, error) {
	data, err := readFile(lang, cat)
	if err != nil {
		return nil, err
	}

	if _, ok := samplesCache[lang]; !ok {
		samplesCache[lang] = make(map[string][]string)
	}

	samples := strings.Split(strings.TrimSpace(string(data)), "\n")

	samplesCache[lang][cat] = samples
	return samples, nil
}

func readFile(lang, cat string) ([]byte, error) {
	fullpath := fmt.Sprintf("/data/%s/%s", lang, cat)
	file, err := FS(useLocalData).Open(fullpath)
	if err != nil {
		return nil, ErrNoSamplesFn(lang)
	}
	defer file.Close()

	return ioutil.ReadAll(file)
}
