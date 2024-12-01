package lcache

import (
	"fmt"
	"log"
	"testing"
)

var db = map[string]string{
	"Tom":  "630",
	"Jack": "589",
	"Sam":  "567",
}

func TestGet(t *testing.T) {
	loadCounts := make(map[string]int, len(db))
	cacheInfo := NewGroup("scores", 2<<10, GetterFunc(
		func(key string) ([]byte, error) {
			log.Println("[SlowDB] search key", key)
			if v, ok := db[key]; ok {
				if _, ok := loadCounts[key]; !ok {
					loadCounts[key] = 0
				}
				loadCounts[key] += 1
				return []byte(v), nil
			}
			return nil, fmt.Errorf("key not found, %s", key)
		}))
	for k, v := range db {
		if view, err := cacheInfo.Get(k); err != nil || view.String() != v {
			t.Fatal("failed to get", k, "from lcache", view)
		}
		if _, err := cacheInfo.Get("unknown"); err == nil {
			t.Fatal("failed to get", k, "from lcache")
		}
	}

}
