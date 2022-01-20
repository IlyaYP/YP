package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
)

func generateRandom(size int) ([]byte, error) {
	b := make([]byte, size)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func main() {
	run1()
	run2()
}

/*
Зашифруем и расшифруем текст с помощью алгоритма AES. Это блочный алгоритм, размер блока — 16 байт.
*/

func run1() {
	src := []byte("12345678901234567") // данные, которые хотим зашифровать
	fmt.Printf("original: %s\n", src)

	// константа aes.BlockSize определяет размер блока и равна 16 байтам
	key, err := generateRandom(aes.BlockSize) // ключ шифрования
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	// получаем cipher.Block
	aesblock, err := aes.NewCipher(key)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	dst := make([]byte, aes.BlockSize) // зашифровываем
	aesblock.Encrypt(dst, src)
	fmt.Printf("encrypted: %x\n", dst)

	src2 := make([]byte, aes.BlockSize) // расшифровываем
	aesblock.Decrypt(src2, dst)
	fmt.Printf("decrypted: %s\n", src2)
}

/*
Чтобы шифровать данные произвольной длины, нужен алгоритм, который делил бы данные на блоки,
преобразовывал и подавал их на вход AES. Стоит взять алгоритм GCM.
Для работы алгоритма GCM нужно дополнительно сгенерировать вектор инициализации из 12 байт.
Вектор должен быть уникальным для каждой процедуры шифрования.
Если переиспользовать один и тот же вектор, можно атаковать алгоритм, подавая на вход данные
с разницей в один байт, и по косвенным признакам вычислить ключ шифрования.
*/

func run2() {
	src := []byte("Этюд в розовых тонах bla bla bla bla bla bla bla bla bla") // данные, которые хотим зашифровать
	fmt.Printf("original: %s\n", src)

	// будем использовать AES256, создав ключ длиной 32 байта
	key, err := generateRandom(2 * aes.BlockSize) // ключ шифрования
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	aesblock, err := aes.NewCipher(key)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	aesgcm, err := cipher.NewGCM(aesblock)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	// создаём вектор инициализации
	nonce, err := generateRandom(aesgcm.NonceSize())
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	dst := aesgcm.Seal(nil, nonce, src, nil) // зашифровываем
	fmt.Printf("encrypted: %x\n", dst)

	src2, err := aesgcm.Open(nil, nonce, dst, nil) // расшифровываем
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	fmt.Printf("decrypted: %s\n", src2)
}
