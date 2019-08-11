package main

import "fmt"

/**
 * @Description:
 * @Date: 2019-08-11 21:32
 * @Author: yaoyao
 */
func main() {
	a := 10
	b := 20

	defer func() {
		fmt.Printf("内部：a = %d, b = %d\n", a, b)
	}() //()代表调用此匿名函数

	defer func(a, b int) {
		fmt.Printf("内部传参：a = %d, b = %d\n", a, b)
	}(a, b) //把参数传递过去，已经先传递参数，只是没有调用

	a = 111
	b = 222
	fmt.Printf("外部：a = %d, b = %d\n", a, b)
}

/*
输出：
外部：a = 111, b = 222
内部传参：a = 10, b = 20
内部：a = 111, b = 222
*/
