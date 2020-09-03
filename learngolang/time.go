package main

import (
	"fmt"
	"time"
)

var week time.Duration

func main() {
	t := time.Now()
	fmt.Println(t)                                               //2020-08-13 17:55:54.6841461 +0800 CST m=+0.001025101
	fmt.Printf("%02d.%02d.%04d\n", t.Day(), t.Month(), t.Year()) //13.08.2020

	t = time.Now().UTC()
	fmt.Println(t)          //2020-08-13 09:55:54.6901317 +0000 UTC
	fmt.Println(time.Now()) //2020-08-13 17:55:54.6901317 +0800 CST m=+0.007010701

	//正在计算时间
	week = 60 * 60 * 24 * 7 * 1e9 //精确到纳秒
	week_from_now := t.Add(time.Duration(week))
	fmt.Println(week_from_now) //2020-08-20 09:55:54.6901317 +0000 UTC

	fmt.Println(t.Format(time.RFC822)) //13 Aug 20 09:55 UTC
	fmt.Println(t.Format(time.ANSIC))  //Thu Aug 13 09:55:54 2020

	fmt.Println(t.Format("02 Jan 2006 15:04")) //13 Aug 2020 09:55
	s := t.Format("20060102")
	fmt.Println(t, "=>", s) //2020-08-13 09:55:54.6901317 +0000 UTC => 20200813
}
