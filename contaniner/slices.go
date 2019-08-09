package main

import "fmt"

func main() {
	//slice是对底层array的一个视图
	arr := [...]int{0,1,2,3,4,5,6,7}
	fmt.Println("arr[2:6] =", arr[2:6])
	// 第一个到第六个
	fmt.Println("arr[:6] =", arr[:6])
	// 第二个以后的全部
	fmt.Println("arr[2:] =", arr[2:])
	// : 全部
	fmt.Println("arr[:] =", arr[:])

	//slice是对底层array的一个视图
	s2 := arr[:]
	s1 := arr[2:]
	fmt.Println("updateSilice s1")
	updateSilice(s1)
	fmt.Println(s1)
	fmt.Println(arr)
	fmt.Println("updateSilice s2")
	updateSilice(s2)
	fmt.Println(s2)
	fmt.Println(arr)
	ext()
}

/**
[]int 不加长度就是slice
 */
func updateSilice(s []int)  {
	s[0] = 100
}

func ext()  {
	//slice是对底层array的一个视图(view),slice可以向后扩展，不能向前扩展，但也不能超越底层array的cap(索引)
	arr := [...]int{0,1,2,3,4,5,6,7}
	narr := arr[:]
	s1 := arr[2:6] // 2,3,4,5
	s2 := s1[3:5] // 5,6?
	fmt.Println(s1)
	fmt.Println(s2)
	//添加元素如果超越Cap，系统会重新分配一个更大的一个底层数组
	//必须接收append返回值
	s3 := append(narr, 777)
	fmt.Println(narr)
	fmt.Println(s3)
	fmt.Println(narr)
}