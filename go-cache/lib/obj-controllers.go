package lib

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"time"
)

func GoCache(key, value string){
	c := cache.New(3*time.Minute, 5*time.Minute)
	c.Set(key, value, cache.DefaultExpiration)
	// c.Delete("baz")

	// Read from the cache
	//v, found := c.Get(key)
	_, found := c.Get(key)
	if found == false {
		error.Error(fmt.Errorf("error"))
	}
	defer c.Flush()
	//fmt.Println(v)
	c.Delete(key)
}
