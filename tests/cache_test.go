package tests

import (
	"testing"
	"github.com/vitalvirtue/go-lru-cache/pkg/cache"
)

func TestLRUCache(t *testing.T) {
    lruCache := cache.NewLRUCache(2)

    lruCache.Set("A", "1")
    if value, found := lruCache.Get("A"); !found || value != "1" {
        t.Error("Expected to find value '1' for key 'A'")
    }

    lruCache.Set("B", "2")
    lruCache.Set("C", "3") // Bu durumda "A" çıkarılmalı

    if _, found := lruCache.Get("A"); found {
        t.Error("Expected key 'A' to be evicted")
    }
}