package main

import (
	"errors"
	"lcache"
	"log"
	"net/http"
)

var db = map[string]string{
	"Tom":  "630",
	"Jack": "589",
	"Sam":  "567",
}

func main() {
	lcache.NewGroup("scores", 2<<10, lcache.GetterFunc(
		func(key string) ([]byte, error) {
			log.Println("[SlowDB] search key", key)
			if v, ok := db[key]; ok {
				return []byte(v), nil
			}
			return nil, errors.New("not found")
		}))
	addr := "localhost:8080"
	peers := lcache.NewHTTPPool(addr)
	log.Println("lcache server listening on", addr)
	log.Fatal(http.ListenAndServe(addr, peers))
}
