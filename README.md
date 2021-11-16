![TestKube Logo](https://raw.githubusercontent.com/kubeshop/testkube/main/assets/logo-dark-text-full.png)

# Welcome to TestKube Template Executor

TestKube Template Executor is test executor skeleton for [TestKube](https://testkube.io)
You can use it as basic building blocks for new executor

# What is executor

Executor is nothing more than program wrapped into Docker container which gets json (testube.Execution) OpenAPI based document, and returns stream of json output lines (testkube.ExecutorOutput) - each output line is simply wrapped in this JSON, like in structured logging idea. 


# Issues and enchancements 

Please follow to main TestKube repository for reporting any [issues](https://github.com/kubeshop/testkube/issues) or [discussions](https://github.com/kubeshop/testkube/discussions)

## Implemention in 5 steps:

1. Create new repo on top of this template 

2. Implement your own Runner on top of [runner interface](https://github.com/kubeshop/testkube/blob/main/pkg/runner/interface.go

3. Change Dockerfile - use base image of whatever test framework/library you want to use

4. Build and push dockerfile to some repository

5. Register Executor Custom Resource in your cluster 

```yaml
apiVersion: executor.testkube.io/v1
kind: Executor
metadata:
  name: postman-executor
  namespace: testkube
spec:
  executor_type: job
  image: kubeshop/testkube-template-executor:0.0.1
  types:
  - example/test
  volume_mount_path: /mnt/artifacts-storage
  volume_quantity: 10Gix

```

Set up volumes as in following example if you want to use artifacts storage (can be downloaded later in dashboard or by `kubectl testkube` plugin)


## Architecture

Template executor offers you basic building blocks to write new executor based on testkube 
libraries written in Go programming language, but your're not limited only to go, you can 
write in any other programming language like Rust, Javascript, Java or Clojure.

Only thing you'll need is to follow OpenAPI spec for input `testkube.Execution` 
(passed as first argument in JSON form) and all output should be JSON lines 
with `testkube.ExecutorOutput` spec there should be also somewhere final 
`ExecutorOutput` with `ExecutionResult` attached after successful (or failed) test execution.

Resources: 
- [OpenaAPI spec details](https://kubeshop.github.io/testkube/openapi/)
- [Spec in YAML file](https://raw.githubusercontent.com/kubeshop/testkube/main/api/v1/testkube.yaml)

Go based resources for input and output objects:
- input: [`testkube.Execution`](https://github.com/kubeshop/testkube/blob/main/pkg/api/v1/testkube/model_execution.go)
- output line: [`testkube.ExecutorOutput`](https://github.com/kubeshop/testkube/blob/main/pkg/api/v1/testkube/model_executor_output.go)


## Examples

- this template repo - is the simplest one 
- [Postman executor](https://github.com/kubeshop/testkube-executor-postman)
- [Cypress executor](https://github.com/kubeshop/testkube-executor-cypress)
- [Curl executor](https://github.com/kubeshop/testkube-executor-curl)

