package merge

import "testing"

type TestFieldsByNameModel struct {
	ID      int
	Name    string
	Number  float64
	Boolean bool
	Test    string
}

type TestFieldsByNameDTO struct {
	Name    string
	Number  float64
	Boolean bool
	Test    int
}

func TestFieldsByName(t *testing.T) {

	dest := TestFieldsByNameModel{1, "Test", 1.1, true, "a string"}

	src := TestFieldsByNameDTO{"Updated", 1.2, false, 1}

	FieldsByName(&dest, src)

	if dest.Name != "Updated" {
		t.Error("Expected dest.Name to equal Updated, got", dest.Name)
	}

	if dest.Number != 1.2 {
		t.Error("Expected dest.Number to equal 1.2, got", dest.Number)
	}

	if dest.Boolean != false {
		t.Error("Expected dest.Boolean to equal false, got", dest.Boolean)
	}

	if dest.Test != "a string" {
		t.Error("Expected dest.Test to be unchanged due to type missmatch, got", dest.Test)
	}
}
