package main

func main() {
	commands := []CommandsDefinition{
		{
			"name":  "base.print",
			"value": "hello",
		},
		{
			"name":  "base.print",
			"value": "world",
		},
	}
	Execute(commands)
}
