package api

var (
	Data = make(map[string]string)
)

func Add(key, value string) error {
	Data[key] = value
	return nil
}

func Get(key string) (string, error) {

	return "", nil
}