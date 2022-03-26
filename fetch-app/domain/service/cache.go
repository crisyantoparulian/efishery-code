package service

import (
	"time"

	"github.com/OrlovEvgeny/go-mcache"
)

var MCache = mcache.New()

func SetCache(key string, data interface{}, duration time.Duration) error {
	err := MCache.Set(key, data, duration)
	if err != nil {
		return err
	}
	return nil
}

func GetCache(key string) interface{} {
	if data, ok := MCache.Get(key); ok {
		return data
	}
	return nil
}
