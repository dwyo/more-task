package cache

import (
	"testing"
	"time"
)

func TestSet(t *testing.T) {
	mc := NewMemoryCache().(*memCache)
	mc.SetMaxMemory("1KB") // Set max memory to 1KB

	// Test normal behavior
	mc.Set("key1", "value1", 10)
	if val := mc.Get("key1"); val != "value1" {
		t.Errorf("Expected 'value1', got %v", val)
	}

	// Test expiration
	mc.Set("key2", "value2", 1) // Expires in 1 second
	time.Sleep(2 * time.Second) // Wait for expiration
	if val := mc.Get("key2"); val != nil {
		t.Errorf("Expected nil (expired), got %v", val)
	}

	// Test memory limit
	largeValue := make([]int, 1024, 1024) // 1KB value

	mc.Set("key3", largeValue, 10)
	if mc.CurUseMemory > mc.MaxMemorySize {
		t.Errorf("Memory limit exceeded: %d > %d", mc.CurUseMemory, mc.MaxMemorySize)
	}

	// Test edge cases
	mc.Set("", "emptyKey", 10) // Empty key
	if val := mc.Get(""); val != "emptyKey" {
		t.Errorf("Expected 'emptyKey', got %v", val)
	}

	mc.Set("key4", nil, -1) // Negative expiration time
	if val := mc.Get("key4"); val != nil {
		t.Errorf("Expected nil (invalid expiration), got %v", val)
	}
}
