package tests

import (
	"go_utilities_toolkit/utils/redisutils"
	"testing"
)

func TestRedisSetGet(t *testing.T) {
	err := redisutils.RedisSet("testkey", "testvalue", 10)
	if err != nil {
		t.Errorf("Redis set failed: %v", err)
	}
	val, err := redisutils.RedisGet("testkey")
	if err != nil {
		t.Errorf("Redis get failed: %v", err)
	}
	if val != "testvalue" {
		t.Errorf("Expected 'testvalue', got '%s'", val)
	}
}
