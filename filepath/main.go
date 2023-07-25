package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

// Go中的filepath包为文件路径提供了跨平台的实用工具。
// 比如linux和windows下的文件路径分隔符是不一样的，linux下是/，windows下是\。
func main() {
	// Join(elem ...string) string，将多个路径元素拼接成一个路径，会根据需要添加路径分隔符。
	// Join将任意数量的路径元素放入一个单一路径里，会根据需要添加路径分隔符。
	// Join还会删除多余的分隔符和目录，以及处理.和..元素。
	p := filepath.Join("dir1", "subdir1", "filename")
	fmt.Println("p:", p)

	// 使用Join代替手动拼接路径可以避免平台差异带来的问题。
	fmt.Println(filepath.Join("dir1//", "filename"))
	// Join还会删除多余的分隔符和目录，以及处理.和..元素，使路径更加干净。
	fmt.Println(filepath.Join("dir1/../dir1", "filename"))

	// Dir(path string) string，返回路径除去最后一个路径元素的部分，即该路径最后一个元素所在的目录。
	fmt.Println("Dir(p):", filepath.Dir(p))
	// Base(path string) string，返回路径的最后一个元素。
	fmt.Println("Base(p):", filepath.Base(p))

	// 判断路径是否是绝对路径 IsAbs(path string) bool
	fmt.Println(filepath.IsAbs("dir/file"))
	fmt.Println(filepath.IsAbs("/dir/file"))

	filename := "config.json"

	// Ext(path string) string，返回路径文件的扩展名。
	ext := filepath.Ext(filename)
	fmt.Println(ext)

	// 获取文件名，不包含扩展名，使用strings.TrimSuffix(path, suffix string) string
	fmt.Println(strings.TrimSuffix(filename, ext))

	// 寻找两路径之间的相对路径，使用Rel(basepath, targpath string) (string, error)
	// 如果相对路径不存在，会返回错误。
	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

	rel, err = filepath.Rel("a/b", "a/c/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)
}