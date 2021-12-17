/*
Задание 1
В код-сниппете ниже есть два типа: DB — слой обращения к базе данных, Handler — слой бизнес-логики.
Вместо реального запроса в базу данных в код-сниппете стоит заглушка: код просто ждёт секунду и затем отдаёт константу.
Вам нужно дописать строчки кода:

    Получение сигнала отмены контекста.
    Отмена контекста через 500 миллисекунд.
*/
package main

import (
	"context"
	"fmt"
	"time"
)

type DB struct {
}

type User struct {
	Name string
}

func (d *DB) SelectUser(ctx context.Context, email string) (User, error) {
	timer := time.NewTimer(1 * time.Second)
	select {
	case <-timer.C:
		return User{Name: "Gosha"}, nil
	case <-ctx.Done(): // 1 - допишите здесь получение сигнала отмены контекста
		return User{}, fmt.Errorf("context canceled")
	}
}

type Handler struct {
	db *DB
}

type Request struct {
	Email string
}

type Response struct {
	User User
}

func (h *Handler) HandleAPI(ctx context.Context, req Request) (Response, error) {
	u, err := h.db.SelectUser(ctx, req.Email)
	if err != nil {
		return Response{}, err
	}

	return Response{User: u}, nil
}

func main() {
	db := DB{}
	handler := Handler{db: &db}
	ctx, cancel := context.WithCancel(context.Background())

	// 2 - допишите код, который отменяет контекст через 500 миллисекунд
	// time.AfterFunc(500*time.Millisecond, cancel)

	// когда запустите код и он отработает успешно,
	// попробуйте заменить длительность на 2000 миллисекунд
	time.AfterFunc(2000*time.Millisecond, cancel)

	req := Request{Email: "test@yandex.ru"}
	resp, err := handler.HandleAPI(ctx, req)
	fmt.Println(resp, err)
}
