package main

import (
	"encoding/base64"
	"fmt"
)

// Go提供了对base64编码的内建支持，并同时支持标准的和URL兼容的base64编码
func main() {
	
	data := "abc123!?$*&()'-=@~" // 待编码的字符串

	// base64编码实例，需要使用[]byte类型的参数
	sEnc := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)

	// base64解码实例，返回[]byte类型的结果
	sDec, _ := base64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))

	// URL base64编码实例，同样需要使用[]byte类型的参数
	uEnc := base64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)

	// URL base64解码实例，同样返回[]byte类型的结果
	uDec, _ := base64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))

	// 标准base64编码和URL base64编码的编码字符串存在稍许不同（后缀分别为+和-）
	// 但是两者都可以正确解码为原始字符串

}