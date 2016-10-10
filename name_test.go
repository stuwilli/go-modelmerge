package merge

import (
	"testing"

	null "gopkg.in/guregu/null.v3"
)

type TestFieldsByNameModel struct {
	ID      int
	Name    string
	Number  float64
	Boolean bool
	Test    string
	NullInt null.Int
}

type TestFieldsByNameDTO struct {
	Name    string
	Number  float64
	Boolean bool
	Test    int
	NullInt null.Int
}

func TestFieldsByName(t *testing.T) {

	dest := TestFieldsByNameModel{1, "Test", 1.1, true, "a string", null.IntFrom(1)}

	src := TestFieldsByNameDTO{"Updated", 1.2, false, 1, null.IntFrom(2)}

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

	if dest.NullInt.Int64 != 2 {
		t.Error("Expected dest.NullInt to equal 2, got", dest.NullInt.Int64)
	}
}
