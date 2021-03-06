/*
Скопируйте примеры себе в IDE. Создайте файл main_test.go и напишите тесты для каждой функции. 
Важно каждое задание покрыть не одним тестом, а таблицей. 
Например, для функции Abs проверьте тесты на -3, 3, -2.000001, -0.000000003 и другие значения.
*/
package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAbs(t *testing.T) {
	// if val := Abs(-3); val != 3 {
	// 	t.Errorf("Abs expected to be 3; got %v", val)
	// }

	tests := []struct { // добавился слайс тестов
        name   string
        value float64
        want float64
    }{
        {
            name:   "simple test #1", // описывается каждый тест
            value: -3,      // значения, которые будет принимать функция
            want:   3,                // ожидаемое значение
        },
        {
            name:   "simple test #2", // описывается каждый тест
            value: 3,      // значения, которые будет принимать функция
            want:   3,                // ожидаемое значение
        },
        {
            name:   "simple test #3", // описывается каждый тест
            value: -2.000001,      // значения, которые будет принимать функция
            want:   2.000001,                // ожидаемое значение
        },
        {
            name:   "simple test #4", // описывается каждый тест
            value: -0.000000003,      // значения, которые будет принимать функция
            want:   0.000000003,                // ожидаемое значение
        },
        {
            name:   "simple test #5", // описывается каждый тест
            value: -100,      // значения, которые будет принимать функция
            want:   100,                // ожидаемое значение
        },
    }

	// for _, tt := range tests { // цикл по всем тестам
    //     t.Run(tt.name, func(t *testing.T) {
    //         if val := Abs(tt.value); val != tt.want {
    //             t.Errorf("Abs expected to be %v; got %v", tt.want, val)
    //         }
    //     })
    // }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            v := Abs(tt.value)
            assert.Equal(t, tt.want, v) // меняем на функцию Equal из пакета assert
        })
    }

}
