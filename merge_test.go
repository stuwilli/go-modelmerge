package merge

import "testing"

type TestMergetTestSubModel struct {
	Data string
	More string
}

type TestMergeTestModel struct {
	ID          int
	Name        string
	Integer     int
	ZeroInt     int
	UnsignedInt uint
	Boolean     bool
	Float       float64
	Payload     TestMergetTestSubModel
}

func createOriginal() TestMergeTestModel {

	sub := TestMergetTestSubModel{Data: "boom", More: "bap"}
	orig := TestMergeTestModel{
		ID:      1,
		Name:    "Original",
		Payload: sub,
		Integer: 1,
		Float:   1.2,
	}

	return orig
}

func createUpdate() TestMergeTestModel {

	sub := TestMergetTestSubModel{Data: "boom!!!"}
	upd := TestMergeTestModel{
		Name:        "Updated",
		Payload:     sub,
		Integer:     22,
		UnsignedInt: 5,
		ZeroInt:     1,
		Boolean:     true,
		Float:       1.1,
	}

	return upd
}

func TestMerge(t *testing.T) {

	original := createOriginal()
	updates := createUpdate()
	fields := []string{"Name", "Integer", "Payload", "More", "ZeroInt", "UnsignedInt", "Boolean", "Float"}

	SelectedFields(&original, &updates, fields)

	if original.ID != 1 {
		t.Error("Original ID should be 1, got", original.ID)
	}

	if original.Name != "Updated" {
		t.Error("Original Name should have been updated, got", original.Name)
	}

	if original.Integer != 22 {
		t.Error("Original Integer should have been updated to 22, got", original.Integer)
	}

	if original.ZeroInt != 1 {
		t.Error("Original ZeroInt should not be updated, expecting 1, got", original.ZeroInt)
	}

	if original.UnsignedInt != 5 {
		t.Error("Original UnsingedInt should have been updated, expecting 5, got", original.UnsignedInt)
	}

	if original.Payload.More != "bap" {
		t.Error("Original Payload.More should not have been updated, got", original.Payload.More)
	}

	if original.Boolean != true {
		t.Error("Original Boolean should not have been updated, expecting true, got", original.Boolean)
	}

}
