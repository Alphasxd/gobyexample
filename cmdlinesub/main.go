package main

import (
	"flag"
	"fmt"
	"os"
)

// go和git都有很多子命令，比如go build和go get，git commit和git pull等
// flag包也支持子命令，但是flag包本身没有提供子命令的解析，需要我们自己来实现
func main() {
	// 声明一个子命令
	fooCmd := flag.NewFlagSet("foo", flag.ExitOnError)
	// 为子命令定义参数，后面参数代表默认值和说明，返回值是一个指针
	fooEnable := fooCmd.Bool("enable", false, "enable")
	fooName := fooCmd.String("name", "", "name")

	// 声明第二个子命令
	barCmd := flag.NewFlagSet("bar", flag.ExitOnError)
	barLevel := barCmd.Int("level", 0, "level")

	// 子命令应该作为程序的第一个参数传入，所以len(os.Args)应该大于1
	if len(os.Args) < 2 {
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}

	// 通过os.Args切片索引检查哪个子命令被调用了，由于第一个参数是程序路径，所以索引从1开始
	switch os.Args[1] {
	case "foo":
		// 解析子命令的参数，解析成功后，子命令的参数会保存在对应的变量中
		fooCmd.Parse(os.Args[2:])
		// 输出解析后的参数
		fmt.Println("subcommand 'foo'")
		fmt.Println("  enable:", *fooEnable)
		fmt.Println("  name:", *fooName)
		// 输出子命令后的其他参数
		fmt.Println("  tail:", fooCmd.Args())
	case "bar":
		// 解析子命令的参数，解析成功后，子命令的参数会保存在对应的变量中
		barCmd.Parse(os.Args[2:])
		// 输出解析后的参数
		fmt.Println("subcommand 'bar'")
		fmt.Println("  level:", *barLevel)
		// 输出子命令后的其他参数
		fmt.Println("  tail:", barCmd.Args())
	default:
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}
}