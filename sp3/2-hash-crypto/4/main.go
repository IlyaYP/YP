package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

/*
Задание 4
Есть зашифрованное сообщение и пароль. Расшифруйте текст, если известно, что:

    сообщение зашифровано алгоритмами GCM и AES256;
    ключ шифрования получен из пароля хеш-функцией SHA256;
    для вектора инициализации используются последние байты ключа.
*/

const (
	password = "x35k9f"
	msg      = `0ba7cd8c624345451df4710b81d1a349ce401e61bc7eb704ca` +
		`a84a8cde9f9959699f75d0d1075d676f1fe2eb475cf81f62ef` +
		`f701fee6a433cfd289d231440cf549e40b6c13d8843197a95f` +
		`8639911b7ed39a3aec4dfa9d286095c705e1a825b10a9104c6` +
		`be55d1079e6c6167118ac91318fe`
)

func main() {
	// допишите код
	// 1) получите ключ из password, используя sha256.Sum256
	key := sha256.Sum256([]byte(password))
	fmt.Printf("key: %x\n", key)

	// 2) создайте aesblock и aesgcm
	aesblock, err := aes.NewCipher(key[:])
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	aesgcm, err := cipher.NewGCM(aesblock)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	// 3) получите вектор инициализации aesgcm.NonceSize() байт с конца ключа
	nonce := key[len(key)-aesgcm.NonceSize():]

	// 4) декодируйте сообщение msg в двоичный формат
	src, err := hex.DecodeString(msg)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	// 5) расшифруйте и выведите данные

	src2, err := aesgcm.Open(nil, nonce, src, nil) // расшифровываем
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	fmt.Printf("decrypted: %s\n", src2)

}
