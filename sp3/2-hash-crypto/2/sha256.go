package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	src := []byte("Здесь могло быть написано, чем Go лучше Rust. " +
		"Но после хеширования уже не прочитаешь. bla bla bla bla")

	h := sha256.New()
	h.Write(src)
	dst := h.Sum(nil)

	fmt.Printf("%x", dst)
}
