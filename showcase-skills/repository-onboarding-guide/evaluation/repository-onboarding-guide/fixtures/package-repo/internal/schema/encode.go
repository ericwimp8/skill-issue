package schema

import "encoding/json"

//go:generate go run ./generator
func Encode(values []string) []byte {
	encoded, _ := json.Marshal(values)
	return encoded
}

