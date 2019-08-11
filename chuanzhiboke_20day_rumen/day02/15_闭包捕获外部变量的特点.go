package main

import "fmt"

/**
 * @Description:
 * @Date: 2019-08-11 20:23
 * @Author: yaoyao
 */
func main() {

	a := 10
	str := "mike"

	func() {
		//闭包以引用方式捕获外部变量
		a = 666
		str = "go"
		fmt.Printf("内部：a = %d, str = %s\n", a, str) //内部：a = 666, str = go
	}() //()代表直接调用

	fmt.Printf("外部：a = %d, str = %s\n", a, str) //外部：a = 666, str = go

}
