package main

import "fmt"

/**
 * @Description:
 * @Date: 2019-07-31 23:21
 * @Author: yaoyao
 */
func main() {
	var num int
	fmt.Printf("请按下楼层：")
	fmt.Scan(&num)

	switch num { //switch后面写的是变量本身
	case 1:
		fmt.Printf("按下的是1楼\n")
		//break	//go语言保留了break关键字，跳出 switch语言，不写也会默认包含
		fallthrough //不跳出switch语言，后面的无条件执行
	case 2:
		fmt.Printf("按下的是2楼\n")
		//break
		fallthrough
	case 3:
		fmt.Printf("按下的是3楼\n")
		//break
	case 4:
		fmt.Printf("按下的是4楼\n")
		break
	default:
		fmt.Println("按下的是xxx楼\n")
	}
}
