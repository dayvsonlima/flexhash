package flexhash

import (
	"testing"
)

func TestExtract(t *testing.T) {

	result := Extract(struct {
		Name  string
		Email string
	}{
		Name:  "My struct",
		Email: "email@gmail.com",
	}, "Name")

	expect := H{"Name": "My struct"}

	if result["Name"] != expect["Name"] {
		t.Error("Invalid hash")
	}

	if expect["Email"] != nil {
		t.Error("Invalid hash")
	}
}
