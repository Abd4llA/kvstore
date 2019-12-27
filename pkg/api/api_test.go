package api

import "testing"

func TestAdd(t *testing.T) {
	k1 := "k1"
	v1 := "v1"
	error := Add(k1, v1)
	v, ok := Data[k1]
	if !ok {
		t.Errorf("Key was not added %v", k1)
	}
	if v != v1 {
		t.Errorf("Unexpected value %v for key %v", v1, k1)
	}
	if error != nil {
		t.Errorf("Unexpected error %v when adding %v for key %v", error, v1, k1)
	}
}
