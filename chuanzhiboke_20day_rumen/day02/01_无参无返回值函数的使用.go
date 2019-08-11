package main

import "fmt"

/**
 * @Description:
 * @Date: 2019-08-11 10:53
 * @Author: yaoyao
 */
func main() {

	//无参无返回值函数的调用：函数名()
	MyFunc()

}

//无参无返回值函数的定义
func MyFunc() {
	a := 666
	fmt.Println("a = ", a)
}
