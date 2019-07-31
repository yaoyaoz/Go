package main

import "fmt"

/**
 * @Description:
 * @Date: 2019-07-31 22:42
 * @Author: yaoyao
 */
func main() {
	s := "王俊凯"

	if s == "王源" {
		fmt.Println("小可爱")
	} else {
		fmt.Println("小帅比")
	}

	//if支持1个初始化语句，初始化语句和判断条件以分号分隔
	if a := 10; a == 10 {
		fmt.Println("a == 10")
	} else {
		fmt.Println("a != 10")
	}
}
