/*
Задание 1
Конвертируйте строковую константу yamlData в TOML-формат. Выведите результат в консоль.
*/
package main

import (
	"fmt"
	"log"

	"gopkg.in/yaml.v2"
)

type Data struct {
	ID     int    `toml:"id"`
	Name   string `toml:"name"`
	Values []byte `toml:"values"`
}

const yamlData = `
id: 101
name: John Doe
values:
- 11
- 22
- 33
`

func main() {
	// допишите код
	// 1) десериализуйте yamlData в переменную типа Data
	// 2) преобразуйте полученную переменную в TOML
	// 3) выведите в консоль результат

	t := Data{}

	err := yaml.Unmarshal([]byte(yamlData), &t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- t:\n%v\n\n", t)

	// out, err := yaml.Marshal(balance)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(out))
}
