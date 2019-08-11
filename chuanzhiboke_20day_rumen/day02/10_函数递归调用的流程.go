package main

import "fmt"

/**
 * @Description:
 * @Date: 2019-08-11 12:20
 * @Author: yaoyao
 */
func main() {
	test(6)
}

func test(a int) {
	if a == 1 { //函数终止调用条件
		fmt.Println("终止：a = ", a)
		return //终止函数调用
	}
	test(a - 1)
	fmt.Println("a = ", a)
}
