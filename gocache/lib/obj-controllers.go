package lib

import (
	"github.com/allegro/bigcache"
	"github.com/dgraph-io/ristretto"
	"github.com/eko/gocache/cache"
	"github.com/eko/gocache/store"
	"time"
)

func GoCacheBigcache(key, value string){
	bigcacheClient, _ := bigcache.NewBigCache(bigcache.DefaultConfig(5 * time.Minute))
	bigcacheStore := store.NewBigcache(bigcacheClient, nil) // No otions provided (as second argument)
	cacheManager := cache.New(bigcacheStore)
	err := cacheManager.Set(key, []byte(value), nil)

	if err != nil {
		panic(err)
	}

	//var v interface{}
	//if v, err = cacheManager.Get(key); err != nil {
	if _, err = cacheManager.Get(key); err != nil {
		panic(err)
	}

	//fmt.Println(string(v.([]byte)))

	defer func(){
		err := cacheManager.Clear()
		if err != nil {
			panic(err)
		}
	}()
	if err := cacheManager.Delete(key); err != nil {
		panic(err)
	}
}

func GoCacheRistretto(key, value string){
	cache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e7,     // Num keys to track frequency of (10M).
		MaxCost:     1 << 30, // Maximum cost of cache (1GB).
		BufferItems: 64,      // Number of keys per Get buffer.
	})
	if err != nil {
		panic(err)
	}

	cache.Set(key, value, 1) // set a value
	// wait for value to pass through buffers
	time.Sleep(10 * time.Millisecond)

	_, found := cache.Get(key)
	//v, found := cache.Get(key)
	if !found {
		panic("missing value")
	}
	defer cache.Close()
	//fmt.Println(v)
	cache.Del(key)
}
