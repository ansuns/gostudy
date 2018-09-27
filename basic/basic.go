package basic

import (
	"fmt"
	"math"
	"math/cmplx"
)

var (
	aa = 3
	bb = 5
	cc = true
)
func main() {
	varDefault()
	varOfficl()
	varDot()
	euler()
	triangle()
	consts()
	enums()
}

func varDefault()  {
	var a int
	var b string
	fmt.Printf("%d %q\n", a, b)
}

func varOfficl()  {
	var q int = 1
	var e string = "a string"
	var r bool = true
	fmt.Println(q,e,r)
}

func varDot()  {
	zz, x, c, v := 77, 3, false, "vv"
	fmt.Println(zz, x, c, v)
}

//欧拉公式，，一脸蒙逼
func euler()  {
	c := 3 + 4i
	fmt.Println(cmplx.Abs(c))
	//？？？？？？？？？？？？？？？？？？
}

//类型强制转换
func triangle()  {
	var a, b int = 3, 4
	var c int
	//类型强制转换
	c = int(math.Sqrt(float64(a * a + b * b)))
	fmt.Println(c)
}

func consts()  {
	const filename = "abc.txt"
	// const 数字可以当作任意类型使用，不用转换
	const a, b = 3, 4
	fmt.Println(filename, a, b)
}

//枚举类型
func enums()  {
	const (
		cpp = iota
		_
		python
		php
	)

	fmt.Println(cpp, python, php)
}
