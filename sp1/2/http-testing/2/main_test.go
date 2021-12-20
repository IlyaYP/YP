/*
Во втором примере покрыть тестами нужно функцию User.FullName(). 
Для этого также скопируйте код себе в IDE и создайте файл main_test.go.
*/
package main

import (
	"testing"
)

func TestFullName(t *testing.T) {

	tests := []struct { // добавился слайс тестов
        name   string
        value User
        want string
    }{
        {
            name:   "simple test #1", // описывается каждый тест
            value: User{
				FirstName: "Misha",
				LastName:  "Popov",
			},      // значения, которые будет принимать функция
            want:   "Misha Popov",                // ожидаемое значение
        },
        {
            name:   "simple test #2", // описывается каждый тест
            value: User{
				FirstName: "Ilya",
				LastName:  "Suprun",
			},      // значения, которые будет принимать функция
            want:   "Ilya Suprun",                // ожидаемое значение
        },
	}


	for _, tt := range tests { // цикл по всем тестам
        t.Run(tt.name, func(t *testing.T) {
            if val := tt.value.FullName(); val != tt.want {
                t.Errorf("FullName expected to be %v; got %v", tt.want, val)
            }
        })
    }


}