/*
Пример 3
В последнем примере нужно покрыть тестами функцию Family.AddNew(). 
Важно учесть в тестах как положительный сценарий (когда добавляется новый член семьи), 
так и отрицательный (когда возникает ошибка при добавлении).
*/
package main

import (
	"testing"
)

func TestAddNew(t *testing.T) {
	f := Family{}
	tests := []struct { // добавился слайс тестов
        name   string
        p Person
		r Relationship
        want error
    }{
        {
            name:   "simple test #1", // описывается каждый тест
		
            p: Person{
				FirstName: "Misha",
				LastName:  "Popov",
				Age: 56,
			},      // значения, которые будет принимать функция
			r: Father,
            want:   nil,                // ожидаемое значение
        },
        {
            name:   "simple test #2", // описывается каждый тест
            p: Person{
				FirstName: "Drug",
				LastName:  "Mishi",
				Age: 57,
			},      // значения, которые будет принимать функция
			r: Father,
            want:   ErrRelationshipAlreadyExists,                // ожидаемое значение
        },
	}

	for _, tt := range tests { // цикл по всем тестам
        t.Run(tt.name, func(t *testing.T) {
            if val := f.AddNew(tt.r, tt.p) ; val != tt.want {
                t.Errorf("AddNew expected to be %v; got %v", tt.want, val)
            }
        })
    }


}
