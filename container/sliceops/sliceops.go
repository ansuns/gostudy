package main

import "fmt"

func printSlice(arr []int) {
	fmt.Printf("len=%d, cap=%d\n", len(arr), cap(arr))
}

func main() {
	//slice操作
	//添加
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	//定义一个空切片
	var s []int
	fmt.Println(s)
	for i := 0; i < 100; i++ {
		printSlice(s)
		s = append(s, 2*i+1)
	}
	fmt.Println(s)
	s3 := arr[2:6] //从array创建切片
	fmt.Println(s3)
	s4 := append(s3, 100)
	fmt.Println(s4)

	//创建1
	r1 := []int{2, 4, 6, 8}
	fmt.Println(r1)
	//创建2 使用Make函数
	r2 := make([]int, 16)
	//也可以跟上他饿的caps
	r3 := make([]int, 16, 32)
	fmt.Println(r2)
	printSlice(r3)
	fmt.Println(r3)

	//复制
	copy(r2, r1)
	fmt.Println(r2) //[2 4 6 8 0 0 0 0 0 0 0 0 0 0 0 0]

	//删除8和这个元素，就是下标3的删掉
	//r2[:3]获取2 4 6 再加上 r2[4:]就是8后面的0 0 0 0
	//获取下标从4开始的所有元素：加上...
	r2 = append(r2[:3], r2[4:]...)
	fmt.Println(r2)

	//删除头
	r2 = r2[1:] //相当于数组切片？
	fmt.Println(r2)

	//删除尾巴
	r2 = r2[:len(r2)-1]
	fmt.Println(r2)
}
