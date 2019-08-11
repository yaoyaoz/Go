package main

import (
	"calc"
	"fmt"
)

/**
 * @Description:
 * @Date: 2019-08-11 22:56
 * @Author: yaoyao
 */
/*
1、不同目录，包名不一样
2、调用不同包里面的函数，格式：包名.函数名()
3、调用别的包的函数，这个函数名字如果是小写，无法让别人调用，要想别人能调用，必须首字母大写
*/
func main() {
	a := calc.Add(3, 5)
	fmt.Println("a = ", a)
}
