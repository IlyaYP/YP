package main

import (
	"fmt"
  "encoding/json"
)

// Зададим динамический тип MyType, где поле Value статически не задано и тип его значения зависит от значения поля Type:
type (
	// MyType использует динамическую типизацию своих полей
	MyType struct {
		// Type задаёт тип объекта поля Value
		Type int `json:"type"`
		// Value содержит объект типа {TypeA, TypeB}
		Value interface{} `json:"value"`
	}

	// TypeA используется полем Value типа MyType при Type == 1
	TypeA struct {
		StrValue string `json:"str_value"`
	}

	// TypeB используется полем Value типа MyType при Type == 2
	TypeB struct {
		FloatValue float32 `json:"float_value"`
	}
)

// Для типа MyType реализуем интерфейс json.Unmarshaler, добавив метод UnmarshalJSON(data []byte) error. 
// В методе укажем все возможные типы объектов, которые могут содержаться в поле Value:
func (t *MyType) UnmarshalJSON(data []byte) error {
    // чтобы избежать рекурсии при json.Unmarshal, объявляем новый тип
    type MyTypeAlias MyType

    // переопределяем поле внутри анонимной структуры
    aliasValue := &struct {
        *MyTypeAlias
        Value json.RawMessage `json:"value"`
    }{
        // встраиваем указатель на целевой объект для заполнения его полей (кроме Value)
        MyTypeAlias: (*MyTypeAlias)(t),
    }
    if err := json.Unmarshal(data, aliasValue); err != nil {
        return err
    }

    // создаём экземпляр конкретного типа
    switch t.Type {
    case 1:
        t.Value = &TypeA{}
    case 2:
        t.Value = &TypeB{}
    }
    if t.Value != nil {
        // производим Unmarshal json.RawMessage значения
        if err := json.Unmarshal(aliasValue.Value, t.Value); err != nil {
            return fmt.Errorf("value of type (%T): unmarshal: %w", t.Value, err)
        }
    }

    return nil
}


// Произведём JSON-сериализацию и десериализацию экземпляра такого типа и изучим результат:
func main() {
	tOrig := MyType{
		Type: 1,
		Value: TypeA{
			StrValue: "some_string",
		},
	}

	tOrigBz, _ := json.MarshalIndent(tOrig, "", "  ")
	fmt.Printf("tOrigBz JSON:\n%s\n", string(tOrigBz))

	tCopy := MyType{}
	json.Unmarshal(tOrigBz, &tCopy)
	fmt.Printf("tCopy Go: %+v\n", tCopy)
	fmt.Printf("tCopy.Value Go: %+v\n", tCopy.Value)
}
