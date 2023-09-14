package commands

import (
	"go-rpa/internal"
	"go-rpa/pkg/base"
)

type PrintCommand struct {
	Value string
}

func (c PrintCommand) NewCommand(args map[string]any) internal.Command {
	if val, ok := args["value"].(string); ok {
		return PrintCommand{
			Value: val,
		}
	}

	return nil
}

func (c PrintCommand) Execute() {
	printer := base.Printer{}
	printer.Print(c.Value)
}

func (c PrintCommand) GetName() string {
	return "base.print"
}
