package example

import "testing"

func TestNew(t *testing.T) {
	expected := 100

	newStruct := New(expected)
	if newStruct.Value.value != expected {
		t.Fatalf(
			"Unexpected value. Got: %d, Expected: %d",
			newStruct.Value.value,
			expected,
		)
	}
}

func TestNewSlice(t *testing.T) {
	expected := []*Struct{
		{
			Value: &Value{
				value: 0,
			},
		},
	}

	newSlice := NewSlice()

	if newSlice[0].Value.value != expected[0].Value.value {
		t.Fatalf(
			"Unexpected value. Got: %d, Expected: %d",
			newSlice[0].Value.value,
			expected[0].Value.value,
		)
	}
}
