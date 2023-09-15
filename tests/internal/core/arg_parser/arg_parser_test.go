package arg_parser

import (
	"bytes"
	"go-rpa/internal/core/arg_parser"
	"log"
	"os"
	"testing"
)

func TestParseRobotJson(t *testing.T) {
	var stdErr bytes.Buffer
	log.SetOutput(&stdErr)

	tempFile, err := os.CreateTemp("", "testfile.json")
	if err != nil {
		t.Fatal(err)
	}
	defer func(name string) {
		err := os.Remove(name)
		if err != nil {
			t.Fatal(err)
		}
	}(tempFile.Name())

	if _, err := tempFile.Write([]byte(`{"key": "value"}`)); err != nil {
		t.Fatal(err)
	}

	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"cmd", "-f", tempFile.Name()}
	data, err, fatal := arg_parser.ParseRobotJson()
	if err != nil || fatal || string(data) != `{"key": "value"}` {
		t.Errorf("Expected valid output but got error: %v, fatal: %v, data: %s", err, fatal, data)
	}

	os.Args = []string{"cmd"}
	_, _, fatal = arg_parser.ParseRobotJson()
	if !fatal {
		t.Errorf("Expected fatal error due to missing -f option")
	}

	os.Args = []string{"cmd", "-f", "nonexistent.json"}
	_, err, fatal = arg_parser.ParseRobotJson()
	if err == nil || !fatal {
		t.Errorf("Expected error and fatal flag due to invalid file path")
	}
}
