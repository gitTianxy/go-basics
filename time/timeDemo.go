package main

import (
	"time"
	"fmt"
)

const TIME_LAYOUT string = "2006-01-02 15:04:05"

func main()  {
	now := time.Now()
	fmt.Println("***NOW:", now)
	str := time2str(now, TIME_LAYOUT)
	fmt.Println("format:", str)
	ts := time2ts(now)
	fmt.Println("timestamp:", ts)
	fmt.Println("parse", str, ":", parseTime(TIME_LAYOUT, str))
	fmt.Println("time for", ts, ":", time.Unix(ts, 0))
}

func time2str(t time.Time, layout string) string  {
	return t.Format(layout)
}

/**
 * t.Unix(): 精确到秒
 * t.Nanosecond(): 精确到纳秒
 */
func time2ts(t time.Time) int64  {
	return t.Unix()
	//return int64(t.Nanosecond())
}

func parseTime(layout, raw string) time.Time  {
	t, err := time.Parse(layout, raw)
	if err != nil {
		panic("parse " + raw + " error")
	} else {
		return t
	}
}