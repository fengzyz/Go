package main

import (
	fm "fmt"
	"time"
)
var week time.Duration
func main()  {
	t := time.Now();
	fm.Println(t)
	fm.Printf("%02d.%02d.%4d\n", t.Day(), t.Month(), t.Year())

	timeStr := t.Format("2006-01-02 15:04:05")
	fm.Println(timeStr)


	t = time.Now().UTC()
	fm.Println(t) // Wed Dec 21 08:52:14 +0000 UTC 2011
	fm.Println(time.Now()) // Wed Dec 21 09:52:14 +0100 RST 2011
	// calculating times:
	week = 60 * 60 * 24 * 7 * 1e9 // must be in nanosec
	week_from_now := t.Add(week)
	fm.Println(week_from_now) // Wed Dec 28 08:52:14 +0000 UTC 2011
	// formatting times:
	fm.Println(t.Format(time.RFC822)) // 21 Dec 11 0852 UTC
	fm.Println(t.Format(time.ANSIC)) // Wed Dec 21 08:56:34 2011
	fm.Println(t.Format("21 Dec 2011 08:52")) // 21 Dec 2011 08:52
	s := t.Format("20111221")
	fm.Println(t, "=>", s)
}
