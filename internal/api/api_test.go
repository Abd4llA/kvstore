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
		{"Set key and value", "k1", "v1", nil},
		{"Set empty key and value", "", "v2", errors.New(errors.EMPTY_KEY_MSG,
			errors.EMPTY_KEY_CODE)},
		{"Set key and empty value", "k3", "", errors.New(errors.EMPTY_VALUE_MSG,
			errors.EMPTY_VALUE_CODE)},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			error := Set(tc.key, tc.value)
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

func TestGet(t *testing.T) {
	cases := [] struct {
		name	string
		key		string
		expectedValue	string
		expectedError	error
	}{
		{"Get a previously added key value", "k1", "v1", nil},
		{"Get a value for an empty key", "", "", errors.New(errors.EMPTY_KEY_MSG,
			errors.EMPTY_KEY_CODE)},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.key!= "" {
				addError := Set(tc.key, tc.expectedValue)
				if addError != nil {
					t.Errorf("Unexpected error was thrown: %v\n", addError)
				}
			}
			value, getError := Get(tc.key)
			if getError != tc.expectedError {
				t.Errorf("Unexpected error was thrown: %v\nExpected error: %v", getError, tc.expectedError)
			}
			if tc.expectedValue != value {
				t.Errorf("Failed to get the value: %s for the key: %s", tc.expectedValue, tc.key)
			}
		})
	}
}
