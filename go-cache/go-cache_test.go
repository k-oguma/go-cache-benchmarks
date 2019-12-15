package main

import (
	"encoding/json"
	"github.com/k-oguma/go-cache-benchmarks/go-cache/lib"
	"io/ioutil"
	"os"
	"regexp"
	"testing"
)

func BenchmarkAppend_CacheEveryTime(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		lib.GoCache("a", "b")
	}
}

type testData struct {
	Rough struct {
		Lk string `json: "lk"`
		Lv string `json: "lv"`
	}
}

func BenchmarkAppend_CacheForLargeStringsEveryTime(b *testing.B) {
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
		lib.GoCache(testData.Rough.Lk, testData.Rough.Lv)
	}
}
