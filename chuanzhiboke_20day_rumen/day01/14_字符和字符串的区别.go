package main

import "fmt"

/**
 * @Description:
 * @Date: 2019-07-28 10:57.
 * @Author: yaoyao
 */

func main() {
	var ch byte
	var str string

	/*
		字符：
		1、单引号
		2、字符往往都只有一个字符，转义字符除外'\n'
	*/
	ch = 'a'
	fmt.Println("ch = ", ch)

	/*
		字符串：
		1、双引号
		2、字符串由1个或多个字符组成
		3、字符串都是隐藏了一个结束符，'\0'
	*/
	str = "a" //由'a'和'\0'组成了一个字符串
	fmt.Println("str = ", str)

	str = "hello go"
	//只想操作字符串的某个字符，从0开始操作
	fmt.Printf("str[0] = %c, str[1] = %c\n", str[0], str[1])
}
