package main

import "fmt"

/**
 * @Description:
 * @Date: 2019-08-11 10:57
 * @Author: yaoyao
 */
func main() {
	//有参无返回值函数调用：函数名（所需参数）
	MyFunc1(777)

	MyFunc2(12, 34)

	MyFunc3(3, 7)
}

/*
有参无返回值函数的定义，普通参数列表
定义函数时，在函数名后面()定义的参数叫形参
参数传递，只能由实参传递给形参，不能反过来，单向传递
*/
func MyFunc1(a int) {
	fmt.Println("a = ", a)
}

func MyFunc2(a int, b int) {
	fmt.Printf("a = %d, b = %d\n", a, b)
}

func MyFunc3(a, b int) {
	fmt.Printf("a = %d, b = %d\n", a, b)
}
