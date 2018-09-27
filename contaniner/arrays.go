package contaniner

import "fmt"


func main() {
	arr()
	arr2 := [5]int {1, 2, 3, 4, 5}
	fmt.Print("------------------\n")
	myArr(arr2)
}

/**
[]int 和 [5]int是不一样的
[5]int才是数组类型
 */
func myArr(arr [5]int)  {
	for i, v := range arr {
		fmt.Println(i,v)
	}
}

/**
数组是值类型， go语言一般不使用数组，而是使用切片
 */
func arr()  {
	var arr1 [5] int
	arr2 := [3]int {1, 2, 3}
	arr3 := [...]int {2, 4, 6, 8, 10}
	//二维数组 4行5列
	var grid [4][5] int
	fmt.Println(arr1, arr2, arr3, grid)

	//遍历
	fmt.Println("数组遍历：\n")
	for i, v := range arr3 {
		fmt.Println(i, v)
	}
}
