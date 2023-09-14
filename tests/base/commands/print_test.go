package commands

import (
	"go-rpa/internal/base/commands"
	"testing"
)

func TestPrintCommand_NewCommand(t *testing.T) {
	args := map[string]interface{}{"value": "hello"}
	cmd := commands.PrintCommand{}.NewCommand(args)
	if cmd == nil {
		t.Fatalf("Expected a PrintCommand, but got nil")
	}
	printCommand, ok := cmd.(commands.PrintCommand)
	if !ok {
		t.Fatalf("Expected a PrintCommand, but got a different type")
	}
	if printCommand.Value != "hello" {
		t.Errorf("Expected value 'hello', but got %s", printCommand.Value)
	}
}

func TestPrintCommand_NewCommand_InvalidInput(t *testing.T) {
	args := map[string]interface{}{"some_key": 42}
	cmd := commands.PrintCommand{}.NewCommand(args)

	if cmd != nil {
		t.Fatalf("Expected nil, but got a PrintCommand")
	}
}
