package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://.com",
			val: []byte("testdata"),
		},
		{
			key: "https://.com",
			val: []byte("more testdata"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.AddCache(c.key, c.val)
			val, ok := cache.GetCache(c.key)
			if !ok {
				t.Errorf("Expected to find the key")
			}
			if string(val) != string(c.val) {
				t.Errorf("Expected to have matching value")
			}
		})
	}
}
