package main

import (
	"github.com/k-oguma/go-cache-benchmarks/cache/lib"
	"os"
)

func main() {
	lib.GoCache(os.Args[1], os.Args[2])
}

