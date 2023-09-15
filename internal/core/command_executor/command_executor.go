package command_executor

import (
	"fmt"
)

func Execute(commands []map[string]any) {
	for _, command := range commands {
		if IsValidCommandDefinition(command) {
			executeCommand(command)
		} else {
			fmt.Println("Invalid command definition: 'name' key missing or not a string.")
		}
	}
}

func IsValidCommandDefinition(command map[string]any) bool {
	_, ok := command["name"].(string)
	return ok
}

func executeCommand(command map[string]any) {
	name := command["name"].(string)
	if cmd, exists := Commands[name]; exists {
		newCmd := cmd.NewCommand(command)
		newCmd.Execute()
	} else {
		fmt.Printf("Command '%s' does not exist.\n", name)
	}
}
