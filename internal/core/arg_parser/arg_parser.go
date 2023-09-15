// Package arg_parser provides utilities for parsing command-line arguments.
package arg_parser

import (
	"github.com/alexflint/go-arg"
	"log"
	"os"
)

// args holds the command-line arguments for the program.
// It defines a File field that is a required argument,
// specifying the path to a JSON file for robot.
type args struct {
	File string `arg:"-f,--file,required" help:"Path to JSON file"`
}

// ParseRobotJson reads and parses the command-line arguments, then reads the specified file.
// It returns the read file data, any errors encountered, and a boolean indicating if a fatal error occurred.
func ParseRobotJson() ([]byte, error, bool) {
	var args args
	if err := arg.Parse(&args); err != nil {
		log.Println("Error parsing arguments:", err)
		return nil, err, true
	}

	data, err := os.ReadFile(args.File)
	if err != nil {
		log.Println("Error reading file:", err)
		return nil, err, true
	}

	return data, nil, false
}
