/**
1)go语言以包作为管理单位
2)每个文件必须先声明包
3)程序必须有一个main包（重要）
*/
// package hello
package main

import "fmt" //我们需要fmt包中的Print函数

//入口函数
func main1() { //左括号必须和函数名同行
	//打印
	//调用函数，大部分需要导入包
	fmt.Println("hello go") //go语言语句结尾是没有分号的
}
