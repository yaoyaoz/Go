package main

import "fmt"

/**
 * @Description:
 * @Date: 2019-08-11 11:56
 * @Author: yaoyao
 */
func main() {
	//函数调用
	a, b, c := myfunc1()
	fmt.Printf("a = %d, b = %d, c = %d\n", a, b, c)
}

//多个返回值
func myfunc1() (int, int, int) {
	return 1, 2, 3
}

//go官方推荐写法
func myfunc2() (a int, b int, c int) {
	a, b, c = 111, 222, 333
	return
}
