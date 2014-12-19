package faker

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

var (
	ErrInsufficientParams = fmt.Errorf("Insufficient params to lookup the samples")
	ErrNoLanguageFn       = func(lang string) error { return fmt.Errorf("The language passed (%s) is not available", lang) }
	ErrNoSamplesFn        = func(lang string) error { return fmt.Errorf("No samples for language: %s", lang) }
)

// Config used to congigure faker
type Config struct {
	Lang         string
	UseLocalData bool
	EnFallback   bool
}

// Faker is the interface for faking things
type Faker interface {
	FirstName() (string, error)
	LastName() (string, error)
	FullName() (string, error)
}

// cat/subcat/lang/samples
type samplesMap map[string]map[string]map[string][]string

type faker struct {
	lang         string
	useLocalData bool
	enFallback   bool
	samples      samplesMap
}

// NewFaker returns new faker object
func NewFaker(config Config) (Faker, error) {
	if config.Lang == "" {
		config.Lang = "en"
	}

	err := checkLang(config)
	if err != nil {
		return nil, err
	}

	return &faker{
		lang:         config.Lang,
		useLocalData: config.UseLocalData,
		enFallback:   config.EnFallback,
		samples:      make(samplesMap),
	}, nil
}

func checkLang(config Config) error {
	availLangs := []string{"en", "br", "ca", "da", "de", "fi", "fr", "mx", "nl", "se", "sn", "uk", "us", "ja", "ar", "cn", "cs",
		"ga", "it", "kr", "nb", "ph", "th", "vn"}

	found := false
	for _, lang := range availLangs {
		if config.Lang == lang {
			found = true
			break
		}
	}
	if !found {
		return ErrNoLanguageFn(config.Lang)
	}
	return nil
}

func (f *faker) lookup(params ...string) (string, error) {
	if len(params) < 2 {
		return "", ErrInsufficientParams
	}

	cat := params[0]
	subcat := params[1]
	lang := f.lang
	if len(params) > 2 {
		lang = params[2]
	}

	var samples []string
	var err error

	_, ok := f.samples[cat]
	if ok {
		samples, ok = f.samples[cat][subcat][lang]
	}

	if !ok {
		samples, err = f.populateSamples(cat, subcat, lang)
		if err != nil {
			return "", err
		}
	}

	return samples[r.Intn(len(samples))], nil
}

func (f *faker) readSamplesFile(cat, subcat, lang string) ([]byte, error) {
	var suffix string
	if lang != "en" {
		suffix = "_" + lang
	}
	fullpath := fmt.Sprintf("/data/%s%s/%s", cat, suffix, subcat)
	file, err := FS(f.useLocalData).Open(fullpath)
	if err != nil {
		if f.lang != "en" && f.enFallback {
			return f.readSamplesFile(cat, subcat, "en")
		}
		return nil, ErrNoSamplesFn(lang)
	}
	defer file.Close()

	return ioutil.ReadAll(file)
}

func (f *faker) populateSamples(cat, subcat, lang string) ([]string, error) {
	data, err := f.readSamplesFile(cat, subcat, lang)
	if err != nil {
		return nil, err
	}

	if _, ok := f.samples[cat]; !ok {
		f.samples[cat] = make(map[string]map[string][]string)
	}

	if _, ok := f.samples[cat][subcat]; !ok {
		f.samples[cat][subcat] = make(map[string][]string)
	}

	samples := strings.Split(string(data), "\n")

	f.samples[cat][subcat][lang] = samples
	return samples, nil
}
