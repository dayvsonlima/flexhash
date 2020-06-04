package flexhash

import (
	"encoding/json"
)

// H .
type H map[string]interface{}

// Extract .
func Extract(element interface{}, hashKeys ...string) H {

	flexHashElement := ToFlex(element)

	c := make(H)

	for key, value := range flexHashElement {
		if inArray(hashKeys, key) {
			c[key] = value
		}
	}

	return c
}

// ToFlex .
func ToFlex(element interface{}) H {
	a, _ := json.Marshal(element)

	var b H
	json.Unmarshal(a, &b)
	return b
}

func inArray(s []string, field string) bool {
	for _, value := range s {
		if field == value {
			return true
		}
	}

	return false
}
