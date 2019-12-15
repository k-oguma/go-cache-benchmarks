package main

import (
	"github.com/k-oguma/go-cache-benchmarks/gocache/lib"
	"os"
)

func main() {
	lib.GoCacheBigcache(os.Args[1], os.Args[2])
	lib.GoCacheRistretto(os.Args[1], os.Args[2])
}
