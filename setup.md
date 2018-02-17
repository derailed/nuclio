# Running with nuclio

```shell
ku create ns nuclio
ku apply -f nuclio/quay.yml
ku apply -f nuclio/nuclio.yml
ku apply -f nuclio/nuclio-rbac.yml

# Port Forward to the playground 8070
kubectl port-forward -n nuclio $(kubectl get pods -n nuclio -l nuclio.io/app=playground -o jsonpath='{.items[0].metadata.name}') 8070:8070 &
open http://localhost:8070
```

```shell
kubectl create secret docker-registry quay --namespace nuclio \
    --docker-username imhotepio \
    --docker-password fernand~1 \
    --docker-server quay.io \
    --docker-email fernand@imhotepio


---
## Nuctl

### CLI Install

```shell
wget https://github.com/nuclio/nuclio/releases/download/0.2.6/nuctl-0.2.6-darwin-amd64
mv nuctl-0.2.6-darwin-amd64 nuctl
chmod +x nuctl
```

### Running a function

```
nuctl deploy helloworld -n nuclio \
  -p https://raw.githubusercontent.com/nuclio/nuclio/master/hack/examples/golang/helloworld/helloworld.go \
  --registry quay.io/imhotepio
 nuctl invoke -n nuclio helloworld -b yo
 http post `mik ip`:32037 body=yo
```
