package main

import "fmt"

/**
调试数组是值类型
 */
func printArray(arr [5]int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func main() {
	//定义数组，数组是值类型
	var arr1 [5] int        //5个Int
	arr2 := [3]int{1, 2, 3} //冒号等于要赋初值
	arr3 := [...]int{7, 88, 99, 66, 44}
	//二维数组:4行5列
	var grid [4][5]int
	fmt.Println(arr1, arr2, arr3, grid)

	//遍历数组
	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}
	//遍历数组：range关键字
	for i, v := range arr3 {
		fmt.Println(i, v)
	}

	printArray(arr3)
	printArray(arr1)
	//printArray(arr2)
}
