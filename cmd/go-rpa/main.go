package main

import (
	"go-rpa/internal/core/arg_parser"
	"go-rpa/internal/core/command_executor"
	"go-rpa/internal/core/graph_parser"
	"go-rpa/internal/utils"
	"log"
)

func main() {
	log.Println("Starting the application...")

	data, err, done := arg_parser.ParseRobotJson()
	if err != nil {
		log.Fatalf("Error while parsing arguments: %v", err)
		return
	}

	if done {
		log.Println("Exiting the application as done is true.")
		return
	}

	log.Println("Parsing graph...")
	graph, err := graph_parser.ParseGraph(data)
	if err != nil {
		log.Fatalf("Error while parsing graph: %v", err)
		return
	}

	log.Println("Successfully parsed graph.")

	var commands []map[string]any

	log.Println("Converting structs to map...")
	for _, node := range graph {
		commands = append(commands, utils.StructToMap(node))
	}

	log.Println("Successfully converted structs to map.")

	log.Printf("Graph: %v\n", graph)

	log.Println("Executing commands...")
	command_executor.Execute(commands)
	log.Println("Commands executed successfully.")

	log.Println("Exiting the application...")
}
