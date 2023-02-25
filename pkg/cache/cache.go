package cache

import (
	"time"

	"github.com/patrickmn/go-cache"
)

var instance *cache.Cache

func NewCache() *cache.Cache {
	return cache.New(10*time.Minute, 10*time.Minute)
}

func GetCache() *cache.Cache {
	return instance
}
