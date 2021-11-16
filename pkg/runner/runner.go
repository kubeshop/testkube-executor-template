package runner

import (
	"github.com/kubeshop/testkube/pkg/api/v1/testkube"
)

func NewRunner() *ExampleRunner {
	return &ExampleRunner{}
}

// ExampleRunner for template - change me to some valid runner
type ExampleRunner struct {
}

func (r *ExampleRunner) Run(execution testkube.Execution) (testkube.ExecutionResult, error) {
	return testkube.ExecutionResult{
		Status: testkube.StatusPtr(testkube.SUCCESS_ExecutionStatus),
		Output: "exmaple test output",
	}, nil
}
