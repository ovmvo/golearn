package main

import "fmt"

func main()  {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	//新切片
	s1 := a[2:5] //从a[2]开始取3个元素
	s1[1] = 666 //给s1[1]赋值, 底层数组也会发生改变
	fmt.Println("s1 = ", s1)
	fmt.Println("a = ", a)

	//另外新的切片

	s2 := s1[2:7] //s2 的切片可以超过s1的长度, 但是不能超过a的长度

	s2[2] = 777

	fmt.Println("s2 = ", s2)
	fmt.Println("a = ", a)

}
