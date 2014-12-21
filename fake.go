package fake

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

// cat/subcat/lang/samples
type samplesTree map[string]map[string]map[string][]string

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

func (st samplesTree) hasKeyPath(cat, subcat, lang string) bool {
	if _, ok := st[cat]; ok {
		if _, ok = st[cat][subcat]; ok {
			if _, ok = st[cat][subcat][lang]; ok {
				return true
			}
		}
	}
	return false
}

func lookup(cat, subcat, lang string) string {
	var samples []string

	if samplesCache.hasKeyPath(cat, subcat, lang) {
		samples = samplesCache[cat][subcat][lang]
	} else {
		var err error
		samples, err = populateSamples(cat, subcat, lang)
		if err != nil {
			if err.Error() == ErrNoSamplesFn(lang).Error() && lang != "en" && enFallback {
				return lookup(cat, subcat, "en")
			}
			return ""
		}
	}

	return samples[r.Intn(len(samples))]
}

func populateSamples(cat, subcat, lang string) ([]string, error) {
	data, err := readSamplesFile(cat, subcat, lang)
	if err != nil {
		return nil, err
	}

	if _, ok := samplesCache[cat]; !ok {
		samplesCache[cat] = make(map[string]map[string][]string)
	}

	if _, ok := samplesCache[cat][subcat]; !ok {
		samplesCache[cat][subcat] = make(map[string][]string)
	}

	samples := strings.Split(string(data), "\n")

	samplesCache[cat][subcat][lang] = samples
	return samples, nil
}

func readSamplesFile(cat, subcat, lang string) ([]byte, error) {
	fullpath := fmt.Sprintf("/data/%s_%s/%s", cat, lang, subcat)
	file, err := FS(useLocalData).Open(fullpath)
	if err != nil {
		return nil, ErrNoSamplesFn(lang)
	}
	defer file.Close()

	return ioutil.ReadAll(file)
}
