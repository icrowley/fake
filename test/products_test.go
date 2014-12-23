package test

import (
	"testing"

	"github.com/icrowley/fake"
)

func TestProducts(t *testing.T) {
	fake.SetLang("en")

	v := fake.Brand()
	if v == "" {
		t.Error("Brand failed")
	}

	v = fake.ProductName()
	if v == "" {
		t.Error("ProductName failed")
	}

	v = fake.Product()
	if v == "" {
		t.Error("Product failed")
	}

	v = fake.Model()
	if v == "" {
		t.Error("Model failed")
	}
}
