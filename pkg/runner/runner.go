package runner

import (
	"fmt"

	"github.com/kubeshop/testkube/pkg/api/v1/testkube"
	"github.com/kubeshop/testkube/pkg/executor/content"
	"github.com/kubeshop/testkube/pkg/executor/env"
	"github.com/kubeshop/testkube/pkg/executor/output"
	"github.com/kubeshop/testkube/pkg/executor/runner"
	"github.com/kubeshop/testkube/pkg/ui"
)

func NewRunner() *ExampleRunner {
	return &ExampleRunner{
		Fetcher: content.NewFetcher(""),
	}
}

// ExampleRunner for template - change me to some valid runner
type ExampleRunner struct {
	Fetcher content.ContentFetcher
}

func (r *ExampleRunner) Run(execution testkube.Execution) (result testkube.ExecutionResult, err error) {

	// use `execution.Variables` for variables passed from Test/Execution
	// variables of type "secret" will be automatically decoded
	env.NewManager().GetReferenceVars(execution.Variables)
	path, err := r.Fetcher.Fetch(execution.Content)
	if err != nil {
		return result, err
	}

	output.PrintEvent("created content path", path)

	contentType := ""
	if execution.Content.Repository != nil {
		contentType, err = r.Fetcher.CalculateGitContentType(*execution.Content.Repository)
		if err != nil {
			output.PrintLog(fmt.Sprintf("%s Can't detect git conent type: %v", ui.IconCross, err))
			return result, err
		}
	}

	if contentType != string(testkube.TestContentTypeGitDir) {
		output.PrintEvent("using file", execution)
		// TODO implement file based test content for string, git-file, file-uri, git
		//      or remove if not used
	}

	if contentType == string(testkube.TestContentTypeGitDir) {
		output.PrintEvent("using dir", execution)
		// TODO implement file based test content for git-dir, git
		//      or remove if not used
	}

	// TODO run executor here

	// error result should be returned if something is not ok
	// return result.Err(fmt.Errorf("some test execution related error occured"))

	// TODO return ExecutionResult
	return testkube.ExecutionResult{
		Status: testkube.ExecutionStatusPassed,
		Output: "exmaple test output",
	}, nil
}

// GetType returns runner type
func (r *ExampleRunner) GetType() runner.Type {
	return runner.TypeMain
}
