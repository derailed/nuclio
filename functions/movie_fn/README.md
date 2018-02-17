# F1 Sample Nuclio Function

## Steps

```shell
make push
ku apply -f f1.yml
nuctl invoke -n nuclio f1 -b yo
http post `mik ip`:32037 body=yo
```


### Delete

```shell
nuctl del function f1 -n nuclio
```

### Deploy

```shell
nuctl deploy f1 -n nuclio -p ./f1.go -i imhotepio-f1 --registry quay.io/imhotepio --version 0.0.1
```

### Update

```shell
nuctl update f1 -n nuclio -p ./f1.go -i imhotepio-f1 --registry quay.io/imhotepio --version 0.0.1
```

### Run

```shell
nuctl invoke -n nuclio f1 -b yo
```

### List
```code
nuctl get -n nuclio function f1
```

