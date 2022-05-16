package main

import (
	"encoding/base64"
	"fmt"
	"log"
)

/**
Base64是网络上最常见的用于传输8Bit字节代码的编码方式之一，可用于在HTTP环境下传递较长的标识信息。Go 的 encoding/base64 提供了对base64的编解码操作。
*/
func main() {

	str := "www.5lmh.com"
	fmt.Printf("string : %v\n", str)

	input := []byte(str)
	fmt.Printf("[]byte : %v\n", input)

	// 演示base64编码
	encodeString := base64.StdEncoding.EncodeToString(input)
	fmt.Printf("encode base64 : %v\n", encodeString)

	// 对上面的编码结果进行base64解码
	decodeBytes, err := base64.StdEncoding.DecodeString(encodeString)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("decode base64 : %v\n", string(decodeBytes))

	fmt.Println()

	// 如果要用在url中，需要使用URLEncoding
	urlencode := base64.URLEncoding.EncodeToString([]byte(input))
	fmt.Printf("urlencode : %v\n", urlencode)
	//URLEncoding
	urldecode, err := base64.URLEncoding.DecodeString(urlencode)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("urldecode : %v\n", string(urldecode))
}
