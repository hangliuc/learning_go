// 数组
// 可以存放多个统一类型的数据
package main

import "fmt"

func main() {
	var numArr01 [3]int = [3]int{1, 2, 3}
	fmt.Println("numArr01", numArr01)

	var numArr02 = [3]int{5, 6, 7}
	fmt.Println("numArr02", numArr02)

	var numArr03 = [...]int{8, 9, 10}
	fmt.Println("numArr03", numArr03)

	var numArr04 = [...]int{1: 800, 0: 900, 2: 9999}
	fmt.Println("numArr04", numArr04)

	// 遍历
	for i := 0; i < len(numArr01); i++ {
		fmt.Println(i)
	}

	for i, v := range numArr02 {
		fmt.Printf("i=%x v=%v\n", i, v)
	}
	fmt.Println()

	for _, v := range numArr03 {
		fmt.Printf("v=%v\n", v)
	}

	//切片
	//创建
	//var intArr [5]int = [...]int{1, 22, 33, 66, 99}
	//slice := intArr[1:3]
	//fmt.Println("intArr", intArr)
	//fmt.Println("slice 的元素是 =", slice) //前闭后开
	//fmt.Println("slice 的容量是 =", cap(slice))

	var slice []float64 = make([]float64, 5, 10)
	slice[1] = 10
	slice[3] = 20
	fmt.Println(slice)
	fmt.Println("slice的size=", len(slice))
	fmt.Println("slice的cap=", cap(slice))

	slice = append(slice, 30, 40, 50) //切片后追加元素
	fmt.Println(slice)

	//两种方式的区别
	// 方式一直接创建数组，程序员可见
	// 方式二 make创建切片，切片在底层维护
}
