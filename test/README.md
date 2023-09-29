### Golang testing

```bash
# Clean cache
go clean -testcache

# Run all tests
go test ./test/...

# Run specific test using name and location
go test -run TestNewQueryMapR ./test/url
go test -run TestNewQueryMapL ./test/url

# Run all benchmarks
go test -bench=. -benchtime=1000x -benchmem ./test/benchmark/...

# Run specific benchmark
go test -bench=^BenchmarkNewQueryMap -benchtime=1000x -benchmem ./test/benchmark
```
