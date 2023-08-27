// 常量
package main

import "fmt"

const Pi = 3.1415926

const zero = 0.0

const (
	size int64 = 1024
	eof        = -1
)

const a, b, c = 3, 4, "foo"

// 预定义常量 true false iota自增常量

const (
	a0 = iota
	a1 = iota
	a2 = iota
)

const (
	u = iota * 42
	v = iota * 42
	w = iota * 42
)

// 枚举 : 一系列相关的常量
const (
	Sunday = iota
	Monday
	Tueday
	Wednesday
	Thursday
	Firday
	Saturday
	numberday
	numberOfDays
)

func main() {
	fmt.Println(size)
}
