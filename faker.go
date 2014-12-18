package faker

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

// Faker is the interface for faking things
type Faker interface {
	FirstName() (string, error)
	LastName() (string, error)
	FullName() (string, error)
}

type samplesMap map[string]map[string][]string

type faker struct {
	lang         string
	useLocalData bool
	samples      samplesMap
}

// NewFaker returns new faker object
func NewFaker(lang string, useLocalData bool) Faker {
	return &faker{
		lang:         lang,
		useLocalData: useLocalData,
		samples:      make(samplesMap),
	}
}

func (f *faker) lookup(cat, subcat string) (string, error) {
	var samples []string
	var err error

	_, ok := f.samples[cat]
	if ok {
		samples, ok = f.samples[cat][subcat]
	}

	if !ok {
		samples, err = f.loadData(cat, subcat)
		if err != nil {
			return "", err
		}
	}

	return samples[r.Intn(len(samples))], nil
}

func (f *faker) loadData(cat, subcat string) ([]string, error) {
	fullpath := fmt.Sprintf("/data/%s%s/%s", cat, f.langSuffix(), subcat)
	file, err := FS(f.useLocalData).Open(fullpath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	if _, ok := f.samples[cat]; !ok {
		f.samples[cat] = make(map[string][]string)
	}

	samples := strings.Split(string(data), "\n")

	f.samples[cat][subcat] = samples
	return samples, nil
}

func (f *faker) langSuffix() string {
	if f.lang == "en" {
		return ""
	}
	return "_" + f.lang
}
