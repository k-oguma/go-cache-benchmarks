package main

import (
	"encoding/json"
	"github.com/k-oguma/go-cache-benchmarks/gocache/lib"
	"io/ioutil"
	"os"
	"regexp"
	"testing"
)

type testData struct {
	Rough struct {
		Lk string `json: "lk"`
		Lv string `json: "lv"`
	}
}

func BenchmarkAppend_CacheEveryTimeForBigcache(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		lib.GoCacheBigcache("a", "b")
	}
}

func BenchmarkAppend_CacheForLargeStringsEveryTimeForBigcache(b *testing.B) {
	b.ResetTimer()

	dir, err := os.Getwd()
	if err != nil {
		b.Fatal(err)
	}

	reg := regexp.MustCompile("(.*/go-cache-benchmarks/).*")
	realpath := reg.ReplaceAllString(dir, "$1")

	r, err := ioutil.ReadFile(realpath + "testdata/large_values.json")
	if err != nil {
		b.Fatal(err)
	}

	var testData testData
	if err := json.Unmarshal(r, &testData); err != nil {
		b.Fatal(err)
	}

	for i := 0; i < b.N; i++ {
		lib.GoCacheBigcache(testData.Rough.Lk, testData.Rough.Lv)
	}
}

func BenchmarkAppend_CacheEveryTimeForRistretto(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		lib.GoCacheRistretto("a", "b")
	}
}

func BenchmarkAppend_CacheForLargeStringsEveryTimeForRistretto(b *testing.B) {
	b.ResetTimer()

	dir, err := os.Getwd()
	if err != nil {
		b.Fatal(err)
	}

	reg := regexp.MustCompile("(.*/go-cache-benchmarks/).*")
	realpath := reg.ReplaceAllString(dir, "$1")

	r, err := ioutil.ReadFile(realpath + "testdata/large_values.json")
	if err != nil {
		b.Fatal(err)
	}

	var testData testData
	if err := json.Unmarshal(r, &testData); err != nil {
		b.Fatal(err)
	}

	for i := 0; i < b.N; i++ {
		lib.GoCacheRistretto(testData.Rough.Lk, testData.Rough.Lv)
	}
}
