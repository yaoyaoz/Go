package main

import "fmt"

/**
 * @Description:
 * @Date: 2019-08-13 23:32
 * @Author: yaoyao
 */
func main() {
	//id1 := 1
	//id2 := 2
	//id3 := 3

	//数组，同一个类型的集合
	var id [5]int

	//操作数组，通过下标，从0开始，到len()-1
	for i := 0; i < len(id); i++ {
		id[i] = i + 1
		fmt.Printf("id[%d] = %d\n", i, id[i])
	}

}
