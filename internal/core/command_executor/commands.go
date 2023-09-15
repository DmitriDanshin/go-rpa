package command_executor

import (
	"go-rpa/internal"
	"go-rpa/internal/base/commands"
)

var printCommand = new(commands.PrintCommand)

var Commands = map[string]internal.Command{
	printCommand.GetName(): printCommand,
}
