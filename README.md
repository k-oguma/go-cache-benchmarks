# benchmarks caches

- [go-cache](https://github.com/patrickmn/go-cache)
- [cache](https://github.com/akyoto/cache)
- [eko/gocache](https://github.com/eko/gocache)
  - Bigcache
  - Ristretto

## Usage

```bash
make test
```

## For example

```bash
% make test
Benchmark target:  /Users/katsuyuki/go/src/github.com/k-oguma/go-cache-benchmarks/cache
goos: darwin
goarch: amd64
pkg: github.com/k-oguma/go-cache-benchmarks/cache
BenchmarkAppend_CacheEveryTime-4                          560948              1919 ns/op             744 B/op         15 allocs/op
BenchmarkAppend_CacheForLargeStringsEveryTime-4           634150              2236 ns/op             744 B/op         15 allocs/op
PASS
ok      github.com/k-oguma/go-cache-benchmarks/cache    2.546s
----------------------
Benchmark target:  /Users/katsuyuki/go/src/github.com/k-oguma/go-cache-benchmarks/go-cache
goos: darwin
goarch: amd64
pkg: github.com/k-oguma/go-cache-benchmarks/go-cache
BenchmarkAppend_CacheEveryTime-4                          379837              3836 ns/op             909 B/op         11 allocs/op
BenchmarkAppend_CacheForLargeStringsEveryTime-4           222919              4854 ns/op             913 B/op         11 allocs/op
PASS
ok      github.com/k-oguma/go-cache-benchmarks/go-cache 2.717s
----------------------
Benchmark target:  /Users/katsuyuki/go/src/github.com/k-oguma/go-cache-benchmarks/gocache
goos: darwin
goarch: amd64
pkg: github.com/k-oguma/go-cache-benchmarks/gocache
BenchmarkAppend_CacheEveryTimeForBigcache-4                           62          39577989 ns/op        337197180 B/op     11301 allocs/op
BenchmarkAppend_CacheForLargeStringsEveryTimeForBigcache-4            36          37570676 ns/op        337200235 B/op     11302 allocs/op
BenchmarkAppend_CacheEveryTimeForRistretto-4                          74          15981556 ns/op        50624497 B/op        548 allocs/op
BenchmarkAppend_CacheForLargeStringsEveryTimeForRistretto-4           76          15482215 ns/op        50625023 B/op        548 allocs/op
PASS
ok      github.com/k-oguma/go-cache-benchmarks/gocache  6.365s
----------------------
```
