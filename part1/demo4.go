// 类型

// 注意类型转换
package main

import "fmt"

var v1 bool

// 整形

var v3 int32
var v2 = 64 // 在这里需要使用显式的类型声明 / v2自动被推导为int类型

func main() {
	v3 = int32(v2) // 需要将 v2 转换为 int32 类型再赋值给 v3
	fmt.Println(v2, v3)
}
