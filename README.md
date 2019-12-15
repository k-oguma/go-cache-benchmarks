# benchmarks caches

 - [go-cache](https://github.com/patrickmn/go-cache)
 - [cache](https://github.com/akyoto/cache)
 
## Usage

```bash
make test
```

## For example

```bash
% make test
Benchmark target: cache
goos: darwin
goarch: amd64
pkg: github.com/k-oguma/go-cache-benchmarks/cache
BenchmarkAppend_CacheEveryTime-12                         378181              3515 ns/op            1409 B/op         18 allocs/op
BenchmarkAppend_CacheForLargeStringsEveryTime-12          622858              3667 ns/op            1406 B/op         17 allocs/op
PASS
ok      github.com/k-oguma/go-cache-benchmarks/cache        5.288s
----------------------
Benchmark target: go-cache
goos: darwin
goarch: amd64
pkg: github.com/k-oguma/go-cache-benchmarks/go-cache
BenchmarkAppend_CacheEveryTime-12                         317542              3480 ns/op             980 B/op         10 allocs/op
BenchmarkAppend_CacheForLargeStringsEveryTime-12          353802              3547 ns/op             922 B/op         10 allocs/op
PASS
ok      github.com/k-oguma/go-cache-benchmarks/go-cache     2.589s
```