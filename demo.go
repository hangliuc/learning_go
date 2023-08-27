package main

import "fmt"

func main() {
	v1 := 10
	v2 := 20

	t := v1 // t = 10
	v1 = v2 // v1 = 20
	v2 = t  // v2 =10

	fmt.Println(v1, v2)

}
