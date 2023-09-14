package main

import "os"

func main() {
	commands := []map[string]any{
		{
			"name":  "base.print",
			"value": "hello",
			"out":   os.Stdout,
		},
		{
			"name":  "base.print",
			"value": "world",
			"out":   os.Stdout,
		},
	}

	for _, command := range commands {
		if name, ok := command["name"].(string); ok {
			command := Commands[name].NewCommand(command)
			command.Execute()
		}
	}
}
