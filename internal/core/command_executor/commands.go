package command_executor

import (
	"go-rpa/internal/base/commands"
)

var printCommand = new(commands.PrintCommand)

var Commands = CommandsImplementation{
	printCommand.GetName(): printCommand,
}
