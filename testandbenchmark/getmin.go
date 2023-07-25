package main

// 待测试的函数，求两个数的最小值
func IntMin(a, b int) int {
    if a < b {
        return a
    }
    return b
}