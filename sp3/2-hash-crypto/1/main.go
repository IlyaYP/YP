package main

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func main() {
	// определяем слайс нужной длины
	b := make([]byte, 16)
	_, err := rand.Read(b) // записываем байты в массив b
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	fmt.Println(hex.EncodeToString(b))
	str64, err := RndArrB64(16)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	fmt.Println(str64)
}

func RndArrB64(size int) (string, error) {
	// определяем слайс нужной длины
	b := make([]byte, size)
	_, err := rand.Read(b) // записываем байты в массив b
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return "", err
	}
	return fmt.Sprint(base64.StdEncoding.EncodeToString(b)), nil
}
