package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

//函数外定义必须要关键字
//函数外定义是包内部，不是全局变量
var aa = 3
var bb = "aa"
var cc = true

//简单点，放到括号里
var (
	ss = 1
	tt = "pp"
)

func variableZeroValue() {
	//定义变量
	var a int
	var s string
	//fmt.Println(a, s)
	//printf可以跟格式
	fmt.Printf("%d %q\n", a, s)
}

func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

//这样定义可以推断type
func variableTypeDeduction() {
	var a, b, c, s = 1, 2, true, "fff"
	fmt.Println(a, b, c, s)
}

//更加简单的写法:省略var,用:=（冒号等于，只能在函数内使用）
func variableShorter() {
	a, b, c, s := 1, 2, true, "fff"
	fmt.Println(a, b, c, s)
}

//欧拉公式，复数
func euler() {
	fmt.Println(
		cmplx.Pow(math.E, 1i*math.Pi) + 1)
}

//类型转换是强制的
/*func triangle() {
	var a, b int = 3, 4
	var c int
	//显式转换int()
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}*/

func triangle() {
	var a, b int = 3, 4
	fmt.Println(calcTriangle(a, b))
}

func calcTriangle(a, b int) int {
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	return c
}

//定义常量
func consts() {
	const (
		filenam = "abc"
		a, b    = 3, 4
	)
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(c)
}

//枚举
func enums() {
	const (
		cpp = iota //iota 表示这组const是自增类的
		java
		python
		golang
		_ //null
		php
	)

	fmt.Println(cpp, java, python, golang, php)
}

//主函数
func main() {
	fmt.Println("Hello World")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	euler()
	consts()
	enums()
}
