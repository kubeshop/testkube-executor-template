package runner

import (
	"github.com/kubeshop/kubtest/pkg/api/kubtest"
)

func NewRunner() *ExampleRunner {
	return &ExampleRunner{}
}

// ExampleRunner for template - change me to some valid runner
type ExampleRunner struct {
}

func (r *ExampleRunner) Run(execution kubtest.Execution) kubtest.ExecutionResult {
	return kubtest.ExecutionResult{
		Status:    kubtest.ExecutionStatusSuceess,
		RawOutput: "exmaple test output",
	}
}
