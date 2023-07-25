package main

import (
    "fmt"
    "testing"
)

// *testing.T 是传入的测试对象，t.Errorf() 用于输出错误信息
func TestIntMinBasic(t *testing.T) {
    ans := IntMin(2, -2)
    if ans != -2 {
		// t.Error*() 会报告测试失败，但是继续执行测试函数
		// t.Fatal*() 会报告测试失败，并且停止执行测试函数
        t.Errorf("IntMin(2, -2) = %d; want -2", ans)
    }
}

// 为了可以重复单元测试，我们可以使用表驱动测试，
// 使用一个结构体数组，每个结构体包含一个测试用例
// 通过遍历结构体数组，来执行每个测试用例
func TestIntMinTableDriven(t *testing.T) {
    var tests = []struct {
        a, b int
        want int
    }{
        {0, 1, 0},
        {1, 0, 0},
        {2, -2, -2},
        {0, -1, -1},
        {-1, 0, -1},
    }

	// 循环遍历结构体数组，执行每个测试用例
    for _, tt := range tests {
        testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
		// t.Run() 用于执行子测试，第一个参数是子测试的名字，第二个参数是一个匿名函数
        t.Run(testname, func(t *testing.T) {
            ans := IntMin(tt.a, tt.b)
            if ans != tt.want {
                t.Errorf("got %d, want %d", ans, tt.want)
            }
        })
    }
}

// 以Benchmark开头的基准测试通常用于测试一段代码的性能
func BenchmarkIntMin(b *testing.B) {
	var result int
	// 在基准测试中，需要重复执行代码 b.N 次，而不是像测试用例中那样只执行一次
	// b.N 的值是系统根据实际情况去调整的，因此基准测试代码不应该关心 N 的具体值
	for i := 0; i < b.N; i++ {
		result = IntMin(1, 2)
	}
	_ = result
}