package main

import (
	"fmt"
	"github.com/vitalvirtue/go-lru-cache/pkg/cache"
	"github.com/vitalvirtue/go-lru-cache/internal/config"
	"github.com/vitalvirtue/go-lru-cache/internal/utils"
)

func main() {
		cfg := config.NewConfig()
    lruCache := cache.NewLRUCache(cfg.CacheCapacity)
    lruCache.Set("A", "1")
    utils.Log("Added key A with value 1")
    lruCache.Set("B", "2")
    utils.Log("Added key B with value 2")
    
    if value, found := lruCache.Get("A"); found {
        fmt.Println("Cache 'A':", value)
    }
    
    lruCache.Set("C", "3") // B çıkarılmalı
    utils.Log("Added key C with value 3")
    
    if _, found := lruCache.Get("B"); !found {
        fmt.Println("Cache 'B' has been evicted.")
    }
}