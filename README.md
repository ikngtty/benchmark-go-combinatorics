# Benchmark Go Combinatorics

## How to run the benchmark

### Local

```shell
go test -bench=. -benchmem ./...
```

### Docker

```shell
docker run -i --rm -v $(pwd):/go/src/github.com/ikngtty/benchmark-go-combinatorics golang go test -bench=. -benchmem github.com/ikngtty/benchmark-go-combinatorics/combinatorics
```
