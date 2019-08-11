package main

import "fmt"

/**
 * @Description:
 * @Date: 2019-08-11 15:42
 * @Author: yaoyao
 */
func main() {
	var result int
	result = Add(1, 3)
	fmt.Println("result = ", result)

	//声明一个函数类型的变量，变量名叫fTest
	var fTest FuncType
	fTest = Add            //是变量就可以赋值
	result = fTest(10, 20) //等价于Add(10, 20)
	fmt.Println("result2 = ", result)

	fTest = Minus
	result = fTest(10, 6) //等价于Minus(10, 6)
	fmt.Println("result3 = ", result)
}

func Add(a, b int) int {
	return a + b
}

func Minus(a, b int) int {
	return a - b
}

//函数也是一种数据类型，通过type给一个函数类型起名
type FuncType func(int, int) int //没有函数名字，没有{}
