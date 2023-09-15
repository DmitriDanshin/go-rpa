package graph_parser

import (
	"encoding/json"
)

type Node struct {
	Name  string `json:"name"`
	Value any    `json:"value"`
}

func ParseGraph(jsonData []byte) ([]Node, error) {
	var commands []Node
	if err := json.Unmarshal(jsonData, &commands); err != nil {
		return nil, err
	}

	return commands, nil
}
