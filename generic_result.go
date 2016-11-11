package graph

import (
	"encoding/json"
	"fmt"
	"io"
)

func parseGraphResult(rc io.ReadCloser, v interface{}) error {
	dec := json.NewDecoder(rc)
	return dec.Decode(v)
}

type GenericGraphResult map[string]interface{}

func (r *GenericGraphResult) GetString(key string) (string, error) {
	if str, ok := (*r)[key].(string); ok {
		return str, nil
	} else {
		return "", fmt.Errorf("Value for key '%s' is not valid", key)
	}
}
