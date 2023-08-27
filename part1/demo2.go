// 匿名变量
package main

import "fmt"

func Getname() (firstname, lastname, nickname string) {
	return "May", "Chan", "Chibi Maruko"
}

func main() {

	_, _, nickname := Getname()
	fmt.Println(nickname)
}
