package main //把test.go文件归属到main


import (
	"sdk/sdk" 
	"fmt"
)
func main() {
	//输出内容
	fmt.Println(sdk.ReadCard())
	// fmt.Println("xxx")
}