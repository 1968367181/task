package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("hello ")

	now := time.Now()

	nowRight := now.Format("2006-01-02 15:04")
	fmt.Println(nowRight)

	nowWrong := now.Format("2006-01-03日错了 15:05 分错了")
	fmt.Println(nowWrong)

	fmt.Println(now.Format("Mon Jan _2 15:04:05 2006 年"))
	fmt.Println("变态吧！")
}
