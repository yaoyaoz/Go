package main

import "fmt"

/**
 * @Description:
 * @Date: 2019-08-11 11:50
 * @Author: yaoyao
 */
func main() {
	var a int
	a = myfunc01()
	fmt.Println("a = ", a)

	b := myfunc02()
	fmt.Println("b = ", b)

	c := myfunc03()
	fmt.Println("c = ", c)
}

/*
无参有返回值：只有一个返回值
有返回值的函数需要通过return中断函数，通过return返回
*/
func myfunc01() int {
	return 66
}

//给返回值起一个变量名，go推荐写法
func myfunc02() (result int) {
	return 55
}

//给返回值起一个变量名，go推荐写法
func myfunc03() (result int) {
	result = 33
	return result
}
