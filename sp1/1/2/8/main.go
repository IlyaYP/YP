/*
Задание 8
Используя Ticker, напишите программу, которая каждые 2 секунды 10 раз подряд
выводит разницу в секундах между текущим временем и временем запуска программы.
Выводить нужно только целую часть секунд.
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	// timer := time.NewTimer(2 * time.Second) // создаём таймер
	ticker := time.NewTicker(2 * time.Second)
	for i := 0; i < 10; i++ {
		t := <-ticker.C                          // ожидаем срабатывания таймера
		fmt.Println(int(t.Sub(start).Seconds())) // выводим разницу во времени
	}
	ticker.Stop()
}
