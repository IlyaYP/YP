// Testing maps
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

	var mtr1 = struct {
		sync.RWMutex
		mtr map[string]map[string]string
	}{mtr: make(map[string]map[string]string)}
	mtr1.mtr["counter"] = make(map[string]string)
	mtr1.mtr["counter"]["PollCount"] = "1"
	mtr1.mtr["counter"]["PollCount1"] = "1"
	mtr1.mtr["gauge"] = make(map[string]string)
	mtr1.mtr["gauge"]["Alloc"] = "123456"
	fmt.Println(mtr1.mtr)

	v, ok := mtr1.mtr["counter"]
	fmt.Println(v, ok)
	// if _, ok := mtr1.mtr["gauge"]; ok{
	// 	fmt.Println(mtr1.mtr)
	// }
}
func add(m map[string]map[string]int, path, country string) {
	mm, ok := m[path]
	if !ok {
		mm = make(map[string]int)
		m[path] = mm
	}
	mm[country]++
}
