package pkg

import "encoding/json"

func ParseJson[T any](data []byte) (T, error) {
	var result T
	err := json.Unmarshal(data, &result)
	return result, err
}
