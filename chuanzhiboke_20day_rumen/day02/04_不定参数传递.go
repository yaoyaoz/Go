package main

import "fmt"

/**
 * @Description:
 * @Date: 2019-08-11 11:19
 * @Author: yaoyao
 */
func main() {

	test1(1, 2, 3, 4)

}

func test1(args ...int) {
	//全部元素传递给myfunc
	myfunc(args...)

	//只想把后2个参数传递给另外一个函数使用
	//myfunc(args[:2]...) //args[0~2（不包括数字2），传递过去
	myfunc(args[2:]...) //从args[2]开始（包括本身，把后面所有元素传递过去）
}

func myfunc(tmp ...int) {
	for _, data := range tmp {
		fmt.Println("tmp = ", data)
	}
}
