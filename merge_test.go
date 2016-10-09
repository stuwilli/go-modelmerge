package modelmerge

import "testing"

type TestSubModel struct {
	Data string `json:"data"`
}

type TestModel struct {
	ID      int          `json:"id"`
	Name    string       `json:"name"`
	Payload TestSubModel `json:"data"`
}

func createOriginal() TestModel {

	sub := TestSubModel{Data: "boom"}
	orig := TestModel{ID: 1, Name: "Original", Payload: sub}

	return orig
}

func createUpdate() TestModel {

	sub := TestSubModel{Data: "boom!!!"}
	upd := TestModel{Name: "Updated", Payload: sub}

	return upd
}

func TestMerge(t *testing.T) {

	original := createOriginal()
	updates := createUpdate()
	fields := []string{"Name", "Payload"}

	Merge(&original, &updates, fields)

	if original.ID != 1 {
		t.Error("Original ID should be 1, got", original.ID)
	}

	if original.Name != "Updated" {
		t.Error("Original Name should have been updated, got", original.Name)
	}

	if original.Payload.Data != "boom!!!" {
		t.Error("Original Payload should have been updated, got", original.Payload.Data)
	}
}
