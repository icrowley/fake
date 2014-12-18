package faker

import "testing"

func TestFakerInit(t *testing.T) {
	f := &faker{lang: "en", useLocalData: true, samples: make(map[string]map[string][]string)}
	f.loadData("name", "first_names")
	if _, ok := f.samples["name"]; !ok {
		t.Fatal("Loading data does not work")
	}
	if _, ok := f.samples["name"]["first_names"]; !ok {
		t.Fatal("Loading data does not work")
	}
}
