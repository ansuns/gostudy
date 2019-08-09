package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		return a / b, nil
	default:
		return 0, fmt.Errorf("unkonw operation:" + op)
	}
}

/**
dai带余除法。。返回两个值
 */
func div(a, b int) (q, r int) {
	return a / b, a % b
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

/**
函数式编程重写eval()
op函数接受两个参数返回一个Int，apply接受op函数，a,b参数，返回int
 */
func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer() //获得函数的指针
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling fnction %s with args"+"(%d, %d)\n", opName, a, b)
	return op(a, b)
}

/**
可变参数:...
 */
func sum(number ...int) int {
	sum := 0
	for i := range number {
		sum += number[i]
	}
	return sum
}

func swap(a, b *int) {
	*b, *a = *a, *b
}

func swap2(a, b int) (int, int) {
	return b, a
}

func main() {
	if result, err := eval(7, 8, "+"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	q, r := div(13, 3)
	fmt.Println(q, r)
	//这里也可以使用匿名函数
	fmt.Println(apply(pow, 1, 6))

	s := sum(1, 2, 3, 4, 5, 6)
	fmt.Println(s)

	a, b := 3, 4
	swap(&a, &b)
	fmt.Println(a, b)

	x, y := swap2(3, 4)
	fmt.Println(x, y)
}
