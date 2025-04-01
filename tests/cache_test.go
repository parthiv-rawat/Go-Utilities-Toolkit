package tests

import (
	"testing"
	"time"
	"go_utilities_toolkit/utils/cache"
)

func TestCache(t *testing.T) {
	cache.SetCache("key", "value", time.Second*1)
	val, ok := cache.GetCache("key")
	if !ok || val != "value" {
		t.Errorf("Cache retrieval failed")
	}
	time.Sleep(time.Second * 2)
	_, ok = cache.GetCache("key")
	if ok {
		t.Errorf("Expected cache to expire")
	}
}