package api

var (
	Data = make(map[string]string)
)

func Add(key, value string) error {
	Data[key] = value
	return nil
}