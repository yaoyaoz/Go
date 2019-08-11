package main

import "fmt"

/**
 * @Description:
 * @Date: 2019-08-11 12:10
 * @Author: yaoyao
 */
func main() {
	fmt.Println("main func")
	a := funca()
	fmt.Println("main a = ", a)
}

func funca() (a int) {
	a = 111

	b := funcb()
	fmt.Println("funca b = ", b)

	//调用另外一个函数
	fmt.Println("funca a = ", a)
	return //返回
}

func funcb() (b int) {
	b = 222
	fmt.Println("funcb b = ", b)
	return
}
