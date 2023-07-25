package main

import (
    "fmt"
    "time"
)

// Go 支持通过基于描述模板的时间格式化与解析。
func main() {
    p := fmt.Println
    t := time.Now()
	
	// RFC3339是Go语言定义的时间格式化模板
    p(t.Format(time.RFC3339))

	// time.Parse()函数可以解析使用time.Format()函数生成的时间字符串
    t1, e := time.Parse(
        time.RFC3339,
        "2012-11-01T22:08:41+00:00")
	p(e)
    p(t1)

	// Format 和 Parse 使用基于例子的布局来决定日期格式， 一般你只要使用 time 包中提供的布局常量就行了，
	// 但是你也可以实现自定义布局。 布局时间必须使用 Mon Jan 2 15:04:05 MST 2006 的格式， 
	// 来指定 格式化/解析给定时间/字符串 的布局。 时间一定要遵循：2006 为年，15 为小时，Monday 代表星期几等规则。
    p(t.Format("3:04PM"))
    p(t.Format("Mon Jan _2 15:04:05 2006"))
    p(t.Format("2006-01-02T15:04:05.999999-07:00"))
    form := "3 04 PM"
    t2, e := time.Parse(form, "8 41 PM")
	p(e)
    p(t2)

	// 对于纯数字表示的时间，你也可以使用标准的格式化字符串来提出出时间值得组成
    fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
        t.Year(), t.Month(), t.Day(),
        t.Hour(), t.Minute(), t.Second())

	// Parse 函数在输入的时间格式不正确时会返回一个错误
    ansic := "Mon Jan _2 15:04:05 2006"
    _, e = time.Parse(ansic, "8:41PM")
    p(e)
}