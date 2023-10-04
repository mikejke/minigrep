package lib

import (
	"reflect"
	"testing"
)

func TestSearchCaseSensitive(t *testing.T) {
	query := "duct"
	content := `
Golang:
safe, fast, productive.
Pick three.
Duct tape.`

	if !reflect.DeepEqual(Search(query, content), []string{"safe, fast, productive."}) {
		t.Fail()
	}
}

func TestSearchCaseInsensitive(t *testing.T) {
	query := "gO"
	content := `
Golang:
safe, fast, productive.
Pick three.
Go and try it!`

	if !reflect.DeepEqual(SearchCaseInsensitive(query, content), []string{"Golang:", "Go and try it!"}) {
		t.Fail()
	}
}
