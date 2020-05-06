# Docker way (easy way)

```shell
$ docker run algebr/bls-key-agg
```

## docker rebuild example

Update the source code in `main.go`

```shell
$ docker build  -t algebr/bls-key-agg -f Dockerfile .
```

then run it

# Build and run the example

Need to have at least go installed, at least version `1.14`

```shell
$ go build && ./bls-key-agg
```
