# slack-plugin
Argo workflows plugins for slack


### Activate plugins


```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: workflow-controller
spec:
  template:
    spec:
      containers:
        - name: workflow-controller
          env:
            - name: ARGO_EXECUTOR_PLUGINS
              value: "true"

```
### Multi step workflow


```
https://argoproj.github.io/argo-workflows/workflow-inputs/
```

## Naming plugins

The name need to include `-executor-plugin` which will be cut away. Thus when naming things we need to adjust for this. 
```go
strings.TrimSuffix(cm.Name, "-executor-plugin")
```

## Examples

Trigger the api flow

``` bash
curl http://localhost:4355//api/v1/template.execute -d \
'{
  "workflow": {
    "metadata": {
      "name": "my-wf"
    }
  },
  "template": {
    "name": "my-tmpl",
    "inputs": {},
    "outputs": {},
    "plugin": {
      "hello": {}
    }
  }
}'
# ...
HTTP/1.1 200 OK
{
  "node": {
    "phase": "Succeeded",
    "message": "Hello template!"
  }
}

```

```bash
curl http://localhost:4355/api/v1/template.execute -d \
{
  "workflow": {
    "metadata": {
      "name": "my-wf"
    }
  },
  "template": {
    "name": "my-tmpl",
    "inputs": {},
    "outputs": {},
    "plugin": {
      "test": {
        "channel": "C035Q8CELGM",
        "text": "Hello Niklas"
      }
    }
  }
}
```