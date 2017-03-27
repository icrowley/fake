package test

import (
	"github.com/icrowley/fake"

	"testing"
)

func TestSetLang(t *testing.T) {
	err := fake.SetLang("ru")
	if err != nil {
		t.Error("SetLang should successfully set lang")
	}

	err = fake.SetLang("sd")
	if err == nil {
		t.Error("SetLang with nonexistent lang should return error")
	}
}

func TestFakerRuWithoutCallback(t *testing.T) {
	fake.SetLang("ru")
	fake.EnFallback(false)
	brand := fake.Brand()
	if brand != "" {
		t.Error("Fake call with no samples should return blank string")
	}
}

func TestFakerRuWithCallback(t *testing.T) {
	fake.SetLang("ru")
	fake.EnFallback(true)
	brand := fake.Brand()
	if brand == "" {
		t.Error("Fake call for name with no samples with callback should not return blank string")
	}
}

// TestConcurrentSafety runs fake methods in multiple go routines concurrently.
// This test should be run with the race detector enabled.
func TestConcurrentSafety(t *testing.T) {
	workerCount := 10
	doneChan := make(chan struct{})

	for i := 0; i < workerCount; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				fake.FirstName()
				fake.LastName()
				fake.Gender()
				fake.FullName()
				fake.Day()
				fake.Country()
				fake.Company()
				fake.Industry()
				fake.Street()
			}
			doneChan <- struct{}{}
		}()
	}

	for i := 0; i < workerCount; i++ {
		<-doneChan
	}
}
