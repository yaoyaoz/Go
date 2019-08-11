package main

import "fmt"

/**
 * @Description:
 * @Date: 2019-08-11 20:37
 * @Author: yaoyao
 */
func main() {

	fmt.Println("aaaaaaaaaaaaaaaaaa")

	//defer延迟调用，main函数结束前调用
	defer fmt.Println("bbbbbbbbbbbbbbbbbb")

	fmt.Println("cccccccccccccccccc")
}
