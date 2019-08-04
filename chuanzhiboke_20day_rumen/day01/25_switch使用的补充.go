package main

import "fmt"

/**
 * @Description:
 * @Date: 2019-08-04 19:19
 * @Author: yaoyao
 */
func main() {

	//支持一个初始化语句，初始化语句和变量本身，以分号分隔
	switch num := 3; num { //switch后面写的是变量本身
	case 1:
		fmt.Printf("按下的是1楼\n")
		//break	//go语言保留了break关键字，跳出 switch语言，不写也会默认包含
		fallthrough //不跳出switch语言，后面的无条件执行
	case 2:
		fmt.Printf("按下的是2楼\n")
		//break
		fallthrough
	case 3, 7, 8:
		fmt.Printf("按下的是3/7/8楼\n")
		//break
	case 4:
		fmt.Printf("按下的是4楼\n")
		break
	default:
		fmt.Println("按下的是xxx楼\n")
	}

	score := 85
	switch { //可以没有条件
	case score > 90: //case后面可以放条件
		fmt.Println("优秀")
	case score > 80:
		fmt.Println("良好")
	case score > 70:
		fmt.Println("一般")
	default:
		fmt.Println("其他")
	}
}
