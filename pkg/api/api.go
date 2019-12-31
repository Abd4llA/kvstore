package api

import "kvstore/pkg/errors"

var (
	Data = make(map[string]string)
)

func Add(key, value string) error {
	if key == "" {
		return errors.New(errors.EMPTYKEYMSG, errors.EMPTYKEYCODE)
	}
	if value == "" {
		return errors.New(errors.EMPTYVALUEMSG, errors.EMPTYVALUECODE)
	}
	Data[key] = value
	return nil
}

func Get(key string) (string, error) {
	if key == "" {
		return "", errors.New(errors.EMPTYKEYMSG, errors.EMPTYKEYCODE)
	}
	return Data[key], nil
}