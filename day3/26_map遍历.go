package main

import (
	"fmt"
)

func main()  {
	m := map[int]string{1: "mike", 2:"yoyo", 3:"go"}

	//第一个返回值位key, 第二个位value, 遍历结果是无序的
	for key,value := range m {
		fmt.Printf("m[%d] = %s\n",key ,value )
	}

	// 如何判断一个key值是否存在
	// 第一个返回值为key所对应的value, 第二个返回值为key是否存在的条件, 存在时ok的值为true
	value , ok := m[4]
	if ok == true {
		fmt.Println("m[1] = ", value)
	} else {
		fmt.Println("key不存在")
	}
}
