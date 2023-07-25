package main

import (
    "fmt"
    "time"
)

// 获取Unix时间的秒数和毫秒数
func main() {

    now := time.Now()
	// Unix时间是从1970年1月1日UTC时间开始到现在的秒数
    secs := now.Unix()
	// UnixNano时间是从1970年1月1日UTC时间开始到现在的纳秒数
    nanos := now.UnixNano()

	// 打印当前时间
    fmt.Println(now)

	// UnixMillis不存在，需要手动计算
    millis := nanos / 1000000

	// 打印Unix时间的秒数
    fmt.Println(secs)
	// 打印Unix时间的毫秒数
    fmt.Println(millis)
	// 打印UnixNano时间的纳秒数
    fmt.Println(nanos)

	// 将整数秒数或者纳秒数转换为相应的时间
    fmt.Println(time.Unix(secs, 0))
    fmt.Println(time.Unix(0, nanos))
}