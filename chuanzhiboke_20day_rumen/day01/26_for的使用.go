package main

import "fmt"

/**
 * @Description:
 * @Date: 2019-08-04 20:51
 * @Author: yaoyao
 */
func main() {

	//for 初始化条件；判断条件；条件变化 ｛｝
	sum := 0
	for i := 1; i <= 100; i++ {
		sum = sum + i
	}
	fmt.Println("sum = ", sum)

}
