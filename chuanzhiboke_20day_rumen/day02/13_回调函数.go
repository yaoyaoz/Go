package main

import "fmt"

/**
 * @Description:
 * @Date: 2019-08-11 15:50
 * @Author: yaoyao
 */
func main() {
	a := Calc(1, 3, Add1)
	fmt.Println("a = ", a)

	b := Calc(1, 3, Minus1)
	fmt.Println("b = ", b)

	c := Calc(1, 3, Mul1)
	fmt.Println("c = ", c)
}

/*
回调函数，函数参数是函数类型，这个函数就是回调函数
计算器，可以进行四则运算
多态：多种形态，调用同一个接口，不同的表现，可以实现不同表现，加减乘除
先有想法，后面再实现功能
*/
func Calc(a, b int, fTest FuncType1) (result int) {
	fmt.Println("Calc")
	result = fTest(a, b) //这个函数还没有实现
	return
}

type FuncType1 func(int, int) int

//实现想法
func Add1(a, b int) int {
	return a + b
}

func Minus1(a, b int) int {
	return a - b
}

func Mul1(a, b int) int {
	return a * b
}
