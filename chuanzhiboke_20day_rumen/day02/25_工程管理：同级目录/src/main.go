package main

import "fmt"

/**
 * @Description:
 * @Date: 2019-08-11 22:17
 * @Author: yaoyao
 */
func main() {
	fmt.Println("this is main")

	/*
		同一个目录，调用别的文件的函数，直接调用即可，无需包名引用
		【idea设置：Directory	F:\workspaces\github\Go\chuanzhiboke_20day_rumen\day02\25_工程管理：同级目录\src】
	*/
	Test()
}
