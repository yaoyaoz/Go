package main

import "fmt"

/**
 * @Description:
 * @Date: 2019-07-28 9:58.
 * @Author: yaoyao
 */
func main() {
	var str1 string //声明变量
	str1 = "abc"
	fmt.Println("str1 = ", str1)

	//自动推导类型
	str2 := "mike"
	fmt.Printf("str2 type is %T\n", str2)

	//内建函数，len()可以测字符串的长度，有多少个字符
	fmt.Println("len(str2) = ", len(str2))
}
