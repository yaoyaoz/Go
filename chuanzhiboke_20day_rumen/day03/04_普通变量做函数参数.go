package main

import "fmt"

/**
 * @Description:
 * @Date: 2019-08-13 23:17
 * @Author: yaoyao
 */
func main() {
	a, b := 10, 20

	//通过一个函数交换a和b的内容
	swap(a, b) //变量本身传递，值传递（站在变量角度）
	fmt.Printf("main: a = %d, b = %d\n", a, b)
}

func swap(a, b int) {
	a, b = b, a
	fmt.Printf("swap: a = %d, b = %d\n", a, b)
}

/*
输出：
swap: a = 20, b = 10
main: a = 10, b = 20
*/
