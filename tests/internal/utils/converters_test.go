package utils

import (
	"go-rpa/internal/utils"
	"reflect"
	"testing"
)

func TestStructToMap(t *testing.T) {
	type Person struct {
		Name    string `json:"name"`
		Age     int    `json:"age"`
		IsAlive bool   `json:"isAlive"`
	}

	type Car struct {
		Make  string
		Model string
	}

	tests := []struct {
		input  interface{}
		output map[string]interface{}
	}{
		{
			input: Person{
				Name:    "John",
				Age:     30,
				IsAlive: true,
			},
			output: map[string]interface{}{
				"name":    "John",
				"age":     30,
				"isAlive": true,
			},
		},
		{
			input: Car{
				Make:  "Toyota",
				Model: "Camry",
			},
			output: map[string]interface{}{
				"Make":  "Toyota",
				"Model": "Camry",
			},
		},
	}

	for i, test := range tests {
		result := utils.StructToMap(test.input)
		if !reflect.DeepEqual(result, test.output) {
			t.Errorf("Test case %d failed: got %v, expected %v", i, result, test.output)
		}
	}
}
