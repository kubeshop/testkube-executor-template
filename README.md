```
██   ██ ██    ██ ██████  ████████ ███████ ███████ ████████ 
██  ██  ██    ██ ██   ██    ██    ██      ██         ██    
█████   ██    ██ ██████     ██    █████   ███████    ██    
██  ██  ██    ██ ██   ██    ██    ██           ██    ██    
██   ██  ██████  ██████     ██    ███████ ███████    ██    
                                 /kʌb tɛst/ by Kubeshop
                    EXCUTOR TEMPLATE
```

<!-- try to enable it after snyk resolves https://github.com/snyk/snyk/issues/347

Known vulnerabilities: ![testkube](https://snyk.io/test/github/kubeshop/testkube/badge.svg)
![testkube-operator](https://snyk.io/test/github/kubeshop-operator/testkube/badge.svg)
![helm-charts](https://snyk.io/test/github/kubeshop/helm-charts/badge.svg)
-->
                                                           
# Welcome to TestKube Template Executor

Kubetest Template Executor is test executor for [testkube](https://testkube.io)

# Issues and enchancements 

Please follow to main TestKube repository for reporting any [issues](https://github.com/kubeshop/testkube/issues) or [discussions](https://github.com/kubeshop/testkube/discussions)

## Implemention in 4 steps:

1. Create new repo on top of this template 

2. Implement your own Runner on top of [runner interface](https://github.com/kubeshop/testkube/blob/main/pkg/runner/interface.go

3. Change Dockerfile - use base image of whatever test framework/library you want to use

4. Register Executor Custom Resource in your cluster 

```yaml
apiVersion: executor.testkube.io/v1
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

`testkube` will pass particular test execution based on executor registered types (e.g. "postman/collection' or 'cypress/project')

executor/server package introduce 2 endpoints 
- first for pushing execution to queue and 
- second for getting execution details from queue


TODO add architecture diagrams

## API 

Template executor implements [TestKube OpenAPI for executors](https://kubeshop.github.io/testkube/openapi/#operations-tag-executor) (look at executor tag)
