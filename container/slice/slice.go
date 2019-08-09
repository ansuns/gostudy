package main

import "fmt"

func main() {
	//不加长度就是slice:[]int
	//切片
	//1.先定义一个数组
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	//定义切片
	//s := arr[2:6]
	fmt.Println("arr =", arr)
	fmt.Println("arr[2:6] =", arr[2:6]) //从第三个开始到第六个
	fmt.Println("arr[:6] =", arr[:6])   //从第一个开始
	fmt.Println("arr[2:] =", arr[2:])   //从三个开始到结束
	fmt.Println("arr[:] =", arr[:])     //所有
}
