package main

import (
"fmt"
"time"
)

func main() {
	now := time.Now()
	//timeStr := now.Format("Mon 2 Jan 2006 15:4:5 MST") //Sun 19 Sep 2021 15:42:00 MSK
	timeStr := now.Format(time.UnixDate)
	fmt.Println(timeStr)
}
