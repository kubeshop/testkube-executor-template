package runner

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/kelseyhightower/envconfig"
	"github.com/kubeshop/testkube/pkg/api/v1/testkube"
	"github.com/kubeshop/testkube/pkg/executor"
	"github.com/kubeshop/testkube/pkg/executor/content"
)

type Params struct {
	Datadir string // RUNNER_DATADIR
}

func NewPlaywrightRunner() (*PlaywrightRunner, error) {
	var params Params
	err := envconfig.Process("runner", &params)
	if err != nil {
		return nil, err
	}

	return &PlaywrightRunner{
		Params:  params,
		Fetcher: content.NewFetcher(""),
	}, nil
}

// PlaywrightRunner - implements runner interface used in worker to start test execution
type PlaywrightRunner struct {
	Params     Params
	Fetcher    content.ContentFetcher
	dependency string
}

func (r *PlaywrightRunner) Run(execution testkube.Execution) (result testkube.ExecutionResult, err error) {
	// check that the datadir exists
	_, err = os.Stat(r.Params.Datadir)
	if errors.Is(err, os.ErrNotExist) {
		return result, err
	}

	// if execution.Content.IsFile() {
	// 	output.PrintEvent("using file", execution)
	// 	return result, fmt.Errorf("passing test as single file not implemented yet")
	// }

	// if execution.Content.IsDir() {
	// 	output.PrintEvent("using dir", execution)
	// TODO implement file based test content for git-dir
	//      or remove if not used
	// }

	// TODO run executor here

	// error result should be returned if something is not ok
	// return result.Err(fmt.Errorf("some test execution related error occured"))

	runPath := ""
	if execution.Content.Repository != nil && execution.Content.Repository.WorkingDir != "" {
		runPath = filepath.Join(r.Params.Datadir, "repo", execution.Content.Repository.WorkingDir)
	}

	if _, err := os.Stat(filepath.Join(runPath, "package.json")); err == nil {
		out, err := executor.Run(runPath, "pnpm", nil, "install")
		if err != nil {
			return result, fmt.Errorf("pnpm install error: %w\n\n%s", err, out)
		}
	}

	out, err := executor.Run(runPath, "pnpm dlx playwright test", nil, "")
	if err != nil {
		return result, fmt.Errorf("playwright binary install error: %w\n\n%s", err, out)
	}

	// return ExecutionResult
	return testkube.ExecutionResult{
		Status: testkube.ExecutionStatusPassed,
		Output: string(out[:]),
	}, nil
}
