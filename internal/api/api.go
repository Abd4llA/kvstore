package api

import "github.com/kvstore/internal/errors"

var (
	data = make(map[string]string)
)

func Set(key, value string) error {
	if key == "" {
		return errors.New(errors.EMPTY_KEY_MSG, errors.EMPTY_KEY_CODE)
	}
	if value == "" {
		return errors.New(errors.EMPTY_VALUE_MSG, errors.EMPTY_VALUE_CODE)
	}
	data[key] = value
	return nil
}

func Get(key string) (string, error) {
	if key == "" {
		return "", errors.New(errors.EMPTY_KEY_MSG, errors.EMPTY_KEY_CODE)
	}
	return data[key], nil
}

func Delete(key string) error {
	if key == "" {
		return errors.New(errors.EMPTY_KEY_MSG, errors.EMPTY_KEY_CODE)
	}
	delete(data, key)
	return nil
}