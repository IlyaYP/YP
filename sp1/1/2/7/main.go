/*
Задание 7
Андрей родился 26 ноября 1993 года. Посчитайте количество дней до его 100-летия — относительно сегодняшнего дня.
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	// допишите код здесь
	birthday := time.Date(1993, 11, 26, 0, 0, 0, 0, time.Local)
	anniversary := birthday.Add(time.Hour * (24*365*100 + (100 / 4 * 24)))
	// days := anniversary.Sub(time.Now())
	days := time.Until(anniversary).Truncate(time.Hour*24).Hours() / 24
	fmt.Println(birthday, anniversary, days)
}
