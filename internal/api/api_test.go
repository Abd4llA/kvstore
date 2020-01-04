package api

import (
	"kvstore/internal/errors"
	"os"
	"testing"
)

func setup() {
	println("SETUP±!")
	data = make(map[string]string)
}

func shutdown() {

}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	shutdown()
	os.Exit(code)
}

func TestAdd(t *testing.T) {
	cases := [] struct {
		name     string
		key      string
		value    string
		expected error
	}{
		{"Add key and value", "k1", "v1", nil},
		{"Add empty key and value", "", "v2", errors.New(errors.EMPTY_KEY_MSG, errors.EMPTY_KEY_CODE)},
		{"Add key and empty value", "k3", "", errors.New(errors.EMPTY_VALUE_MSG, errors.EMPTY_VALUE_CODE)},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			error := Add(tc.key, tc.value)
			if error != nil {
				if error != tc.expected {
					t.Errorf("Unexpected error %v when adding %v for key %v", error, tc.value, tc.key)
				}
			} else {
				v, ok := data[tc.key]
				if !ok {
					t.Errorf("Key was not added %v", tc.key)
				}
				if v != tc.value {
					t.Errorf("Unexpected value %v for key %v", tc.value, tc.key)
				}
			}
		})
	}
}

//func TestGet(t *testing.T) {
//	k2 := "k2"
//	v2 := "v2"
//	data[k2] = v2
//	val, error := Get(k2)
//	if val != v2 {
//		t.Errorf("Unexpected value %v for key %v, expected %v", val, k2, v2)
//	}
//	if error != nil {
//		t.Errorf("Unexpected error %v when getting key %v", error, k2)
//	}
//}