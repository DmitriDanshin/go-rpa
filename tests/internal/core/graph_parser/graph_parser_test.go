package graph_parser

import (
	"go-rpa/internal/core/graph_parser"
	"reflect"
	"testing"
)

func TestParseGraph(t *testing.T) {
	jsonData := `
	[
		{
			"name": "base.print",
			"value": "hello"
		},
		{
			"name": "base.print",
			"value": "world"
		}
	]
	`

	expected := []graph_parser.Node{
		{Name: "base.print", Value: "hello"},
		{Name: "base.print", Value: "world"},
	}

	result, err := graph_parser.ParseGraph([]byte(jsonData))
	if err != nil {
		t.Fatalf("ParseGraph returned an error: %s", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("ParseGraph returned %v, expected %v", result, expected)
	}
}
