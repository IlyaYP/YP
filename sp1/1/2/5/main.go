/*
Задание 5
В задании 4 вы получили время из строки. Теперь нужно понять, отличается ли время, которое вам передали, от текущего.
Сравните текущее время с результатом, полученным в задании 4.
*/

package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	currentTimeStr := "2021-09-19T15:59:41+03:00"
	// допишите код здесь
	layout := "2006-1-2T15:4:5-07:00"
	currentTime, err := time.Parse(layout, currentTimeStr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(currentTime)

	now := time.Now()

	fmt.Println("Is", now, "before", currentTime, "? Answer:", now.Before(currentTime))
	fmt.Println("Is", now, "after", currentTime, "? Answer:", now.After(currentTime))
	fmt.Println("Is", now, "equal", currentTime, "? Answer:", now.Equal(currentTime))
}