# Movies

### Deploy

```shell
nuctl deploy movie -n nuclio -p function.go -i imhotepio-movie --registry quay.io/imhotepio --version 0.0.1
```


### Delete

```shell
nuctl del function movie -n nuclio
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

