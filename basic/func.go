package main

import "fmt"

func main() {
	if res, err := evel(5, 10,"x"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}

	q, r := div(10, 3)
	fmt.Println(
		q,
		r,)

	fmt.Println(sum(1,2,3,4,5,6,7,8,9))
	a, b := 3, 4
	a, b  = swapN(a, b)
	fmt.Println(a, b)
}

//返回一个值
//错误则返回错误
func evel(a, b int, op string) (int, error)  {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		//return a /b
		q, _ := div(a, b)
		return q, nil
	default:
		return 0, fmt.Errorf(
			"unsupported operation: %s", op)

	}
}

//返回两个值
//带余除法
func div(a, b int) (q, r int) {
	return a /b, a % b
}

/*
op 函数，接收两个Int参数
apply接收一个op函数，和 a, int 参数
并返回Int
 */
func apply(op func(int, int) int, a, b int) int {
	return  1
}

/*
...可变参数列表，go语言的函数只有这个特效，没有重载
 */
func sum(numbers ...int) int {
	s := 0
	for i :=  range numbers  {
		s += numbers[i]
	}

	return s
}

//指针
/*
* + type 数据类型，就表示指针
 */
/*func swap(a, b *int)  {
	*b,*a = *a, *b
}*/

//交换值，正确的写法，不用指针,不可变量
func swapN(a,b int) (int, int)  {
	return b, a
}