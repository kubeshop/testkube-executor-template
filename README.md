```
██   ██ ██    ██ ██████  ████████ ███████ ███████ ████████ 
██  ██  ██    ██ ██   ██    ██    ██      ██         ██    
█████   ██    ██ ██████     ██    █████   ███████    ██    
██  ██  ██    ██ ██   ██    ██    ██           ██    ██    
██   ██  ██████  ██████     ██    ███████ ███████    ██    
                               /kjuːb tɛst/ by Kubeshop
                    EXCUTOR TEMPLATE
```

<!-- try to enable it after snyk resolves https://github.com/snyk/snyk/issues/347

Known vulnerabilities: ![kubtest](https://snyk.io/test/github/kubeshop/kubtest/badge.svg)
![kubtest-operator](https://snyk.io/test/github/kubeshop-operator/kubtest/badge.svg)
![helm-charts](https://snyk.io/test/github/kubeshop/helm-charts/badge.svg)
-->
                                                           
# Welcome to Kubtest Template Executor

Kubetest Template Executor is test executor for [Kubtest](https://kubtest.io)

# Issues and enchancements 

Please follow to main kubtest repository for reporting any [issues](https://github.com/kubeshop/kubtest/issues) or [discussions](https://github.com/kubeshop/kubtest/discussions)

## Implemention in 4 steps:

1. Create new repo on top of this template 

2. Implement your own Runner on top of [runner interface](https://github.com/kubeshop/kubtest/blob/main/pkg/runner/interface.go

3. Change Dockerfile - use base image of whatever test framework/library you want to use

4. Register Executor Custom Resource in your cluster 

```yaml
apiVersion: executor.kubtest.io/v1
kind: Executor
metadata:
  name: example-executor
spec:
  uri: http://template-executor:8082
  types:
  - example/test
```


## Architecture

Template executor implement work queue on top of API REST endpoints
Work queue implemented on top of MongoDB. You can scale workers easily on many nodes.

`Kubtest` will pass particular test execution based on executor registered types (e.g. "postman/collection' or 'cypress/project')

executor/server package introduce 2 endpoints 
- first for pushing execution to queue and 
- second for getting execution details from queue


TODO add architecture diagrams

## API 

Template executor implements [Kubtest OpenAPI for executors](https://kubeshop.github.io/kubtest/openapi/#operations-tag-executor) (look at executor tag)
