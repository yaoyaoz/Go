package main

import "fmt"

func main() {
	//1、声明变量，没有初始化，零值（初始值）为false
	var a bool
	fmt.Println("a = ", a)

	//2、自动推导类型
	var b = true
	fmt.Printf("b type  is %T\n", b)
	fmt.Println("b = ", b)

	c := false
	fmt.Println("c = ", c)
}
