package main

import (
	"fmt"
	"os"
)

/**
 * @Description:
 * @Date: 2019-08-11 21:40
 * @Author: yaoyao
 */
func main() {
	//接收用户传递的参数，都是以字符串方式传递
	list := os.Args

	n := len(list)
	fmt.Println("n = ", n)

	for i := 0; i < n; i++ {
		fmt.Printf("list[%d] = %s\n", i, list[i])
	}

	fmt.Println("======================")

	for i, data := range list {
		fmt.Printf("list[%d] = %s\n", i, data)
	}
}
