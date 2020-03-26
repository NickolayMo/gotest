package app

import (
	"fmt"
	"testing"
)

type TestProcessor struct {}

func (t *TestProcessor)Process(s string) (string, error) {
	return s, nil
}

type TestProcessorFail struct {}

func (t *TestProcessorFail)Process(s string) (string, error) {
	return "", fmt.Errorf("test err")
}


func TestRunner_RunSuccess(t *testing.T) {
	test := "test"
	runner := NewRunner(1, test, &TestProcessor{})
	runner.Run(test)
	result := runner.GetResults()
	if result != 1 {
		t.Error("Expected 1 got ", result)
	}
}

func TestRunner_RunFail(t *testing.T) {
	test := "test"
	runner := NewRunner(1, test, &TestProcessorFail{})
	runner.Run(test)
	result := runner.GetResults()
	if result != 0 {
		t.Error("Expected 0 got ", result)
	}
}

