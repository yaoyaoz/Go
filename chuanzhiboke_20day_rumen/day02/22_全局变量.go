package main

import "fmt"

/**
 * @Description:
 * @Date: 2019-08-11 21:51
 * @Author: yaoyao
 */

/*
定义在函数外部的变量是全局变量
全局变量在任何地方都能使用
*/
var a = 10

func main() {
	a = 20
	fmt.Println("a = ", a)

	test3()
}

func test3() {
	fmt.Println("test a = ", a)
}
