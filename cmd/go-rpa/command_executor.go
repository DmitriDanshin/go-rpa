package main

import (
	"fmt"
)

func Execute(commands []CommandsDefinition) {
	for _, command := range commands {
		if IsValidCommandDefinition(command) {
			executeCommand(command)
		} else {
			fmt.Println("Invalid command definition: 'name' key missing or not a string.")
		}
	}
}

func IsValidCommandDefinition(command CommandsDefinition) bool {
	_, ok := command["name"].(string)
	return ok
}

func executeCommand(command CommandsDefinition) {
	name := command["name"].(string)
	if cmd, exists := Commands[name]; exists {
		newCmd := cmd.NewCommand(command)
		newCmd.Execute()
	} else {
		fmt.Printf("Command '%s' does not exist.\n", name)
	}
}
