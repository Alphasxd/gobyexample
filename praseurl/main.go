package main

import (
	"fmt"
	"net"
	"net/url"
)

// 使用Go来解析URL，可以使用内置的net/url包
func main() {
	// URL示例，包含了一个scheme，认证信息(userinfo)，主机名和端口(@host)，路径(/path)，查询参数(?query)和片段(#fragment)
	//	[scheme:][//[userinfo@]host][/]path[?query][#fragment]
	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	// 解析这个URL并确保解析没有出错
	u, err := url.Parse(s)
	// 使用recover函数截获panic异常并打印错误信息
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("URL Parse Error:\n", r)
		}
	}()
	if err != nil {
		panic(err)
	}

	// 直接访问scheme
	fmt.Println(u.Scheme)

	// User包含了所有的认证信息，这里调用Username和Password来获取独立值
	fmt.Println(u.User)
	fmt.Println(u.User.Username())
	p, _ := u.User.Password()
	fmt.Println(p)

	// Host包含了主机名和端口，如果端口存在的话，使用strings.Split函数从Host中手动提取端口
	fmt.Println(u.Host)
	_, port, _ := net.SplitHostPort(u.Host) // 这是net包的SplitHostPort函数
	fmt.Printf("net.SplitHostPort: %v\n", port)
	fmt.Printf("url.Port(): %v\n", u.Port()) // 这里使用的是net/url包的Port()函数，由url包的splitHostPort函数调用

	// 这里我们提取路径和查询片段信息
	fmt.Println(u.Path)
	fmt.Println(u.Fragment)

	// 要得到字符串中的k=v这种格式的查询参数，可以使用RawQuery函数
	// 你也可以将查询参数解析为一个map，这个map以查询参数名为键，对应的值字符串切片为值，使用url.ParseQuery函数
	fmt.Println(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	fmt.Println(m["k"][0]) // 只想得到一个键对应的第一个值，可以直接使用索引获取
}