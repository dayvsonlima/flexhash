package flexhash

import (
	"encoding/json"
)

// H .
type H map[string]interface{}

// Extract .
func Extract(element interface{}, hashKeys ...string) H {
	a, _ := json.Marshal(element)

	var b H
	c := make(H)

	json.Unmarshal(a, &b)

	for key, value := range b {
		if inArray(hashKeys, key) {
			c[key] = value
		}
	}

	return c
}

func inArray(s []string, field string) bool {
	for _, value := range s {
		if field == value {
			return true
		}
	}

	return false
}
