package lib

import (
	"fmt"
	"github.com/akyoto/cache"
	"time"
)

func GoCache(key, value string){
	c := cache.New(5 * time.Minute)
	c.Set(key, value, 5 * time.Minute)

	// Read from the cache
	//k, found := c.Get(key)
	_, found := c.Get(key)
	if found == false {
		error.Error(fmt.Errorf("error"))
	}

	defer c.Close()
	//fmt.Println(k)
	c.Delete(key)
}
