package main

import "fmt"

/**
 * @Description:
 * @Date: 2019-07-28 16:10.
 * @Author: yaoyao
 */
func main() {
	var a int //声明变量
	fmt.Printf("请输入变量a：")

	//阻塞等待用户的输入
	//fmt.Scanf("%d", &a)	//别忘了&
	fmt.Scan(&a)
	fmt.Println("a = ", a)
}
