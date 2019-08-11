package main

import "fmt"

/**
 * @Description:
 * @Date: 2019-08-11 20:39
 * @Author: yaoyao
 */
func main() {

	//defer 先进后出的顺序执行，哪怕函数或某个延迟调用发生错误，这些调用依旧会被执行
	defer fmt.Println("aaaaaaaa")

	defer fmt.Println("bbbbbbbb")

	//调用一关函数，导致内存出问题
	defer testE(0)

	defer fmt.Println("cccccccc")

}

/*
输出：
cccccccc
bbbbbbbb
aaaaaaaa
panic: runtime error: integer divide by zero
*/

func testE(x int) {
	fmt.Println(100 / x)
}
