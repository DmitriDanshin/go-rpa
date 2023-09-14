package internal

type Command interface {
	Execute()
	GetName() string
	NewCommand(args map[string]any) Command
}
