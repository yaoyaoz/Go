package main

import (
	"fmt"
	"time"
)

/**
 * @Description:
 * @Date: 2019-08-04 22:38
 * @Author: yaoyao
 */
func main() {

	i := 0
	for { //for后面不写任何东西，这个循环条件永远为真，死循环
		i++
		time.Sleep(time.Second)
		if i == 5 {
			//break	//跳出循环，如果嵌套多个循环，跳出最近的那个内循环
			continue //跳出本次循环，下一次继续
		}
		fmt.Println("i = ", i)
	}

}
