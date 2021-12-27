package main

import (
	"encoding/json"
	"fmt"
	"sync"
)

func main() {
	mtr := map[string]map[string]string{
		"gauge": {"Alloc": "123456789",
			"BuckHashSys":   "",
			"Frees":         "",
			"GCCPUFraction": "",
			"GCSys":         "",
			"HeapAlloc":     "",
			"HeapIdle":      "",
			"HeapInuse":     "",
			"HeapObjects":   "",
			"HeapReleased":  "",
			"HeapSys":       "",
			"LastGC":        "",
			"Lookups":       "",
			"MCacheInuse":   "",
			"MCacheSys":     "",
			"MSpanInuse":    "",
			"MSpanSys":      "",
			"Mallocs":       "",
			"NextGC":        "",
			"NumForcedGC":   "",
			"NumGC":         "",
			"OtherSys":      "",
			"PauseTotalNs":  "",
			"StackInuse":    "",
			"StackSys":      "",
			"Sys":           "",
			"RandomValue":   ""},
		"counter": {"PollCount": ""},
	}

	res1B, _ := json.Marshal(mtr)
	fmt.Println(string(res1B))

	byt := []byte(string(res1B))
	var dat map[string]interface{}
	json.Unmarshal(byt, &dat)
	fmt.Println(dat)

	hits := make(map[string]map[string]int)
	fmt.Println(hits["/doc/"]["au"])
	add(hits, "/doc/", "au")
	fmt.Println(hits["/doc/"]["au"])

	var counter = struct {
		sync.RWMutex
		m map[string]int
	}{m: make(map[string]int)}

	// To read from the counter, take the read lock:
	counter.RLock()
	n := counter.m["some_key"]
	counter.RUnlock()
	fmt.Println("some_key:", n)

	// To write to the counter, take the write lock:
	counter.Lock()
	counter.m["some_key"]++
	counter.Unlock()
}
func add(m map[string]map[string]int, path, country string) {
	mm, ok := m[path]
	if !ok {
		mm = make(map[string]int)
		m[path] = mm
	}
	mm[country]++
}
