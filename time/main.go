package main

import (
    "fmt"
    "time"
)

func main() {
    p := fmt.Println
	// 获取当前时间
    now := time.Now()
    p(now)

	// 通过提供年月日等信息，可以构建一个time对象
    then := time.Date(
        2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
    p(then)

    p(then.Year()) // 年
    p(then.Month())	// 月
    p(then.Day()) // 日
    p(then.Hour()) // 时
    p(then.Minute()) // 分
    p(then.Second()) // 秒
    p(then.Nanosecond()) // 纳秒
    p(then.Location()) // 时区

    p(then.Weekday()) // 星期几

    p(then.Before(now)) // 比较时间，是否在now之前
    p(then.After(now)) // 比较时间，是否在now之后
    p(then.Equal(now)) // 比较时间，是否相等

    diff := now.Sub(then) // 计算时间差
    p(diff)

    p(diff.Hours()) // 时间差的小时
    p(diff.Minutes()) // 时间差的分钟
    p(diff.Seconds()) // 时间差的秒
    p(diff.Nanoseconds()) // 时间差的纳秒

    p(then.Add(diff)) // 时间加上时间差
    p(then.Add(-diff)) // 时间减去时间差
}