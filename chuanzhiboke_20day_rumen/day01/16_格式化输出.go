package main

import "fmt"

/**
 * @Description:
 * @Date: 2019-07-28 15:40.
 * @Author: yaoyao
 */
func main() {
	a := 10
	b := "abc"
	c := 'a'
	d := 3.14
	e := false

	//%T操作变量所属类型
	fmt.Printf("%T, %T, %T, %T, %T\n", a, b, c, d)

	/*
		%d 整型格式
		%s 字符串格式
		%c 字符格式
		%f 浮点型格式
		%t 布尔型格式
	*/
	fmt.Printf("a = %d, b = %s, c = %c, d = %f, e = %t\n", a, b, c, d, e)
	//%v 自动匹配格式输出
	fmt.Printf("a = %v, b = %v, c = %v, d = %v, e = %v\n", a, b, c, d, e)
}
