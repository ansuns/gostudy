package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

//返回的是一个无参数的函数
func fibonacci() intGen {
	//前两个数
	a, b := 0, 1
	return func() int {
		//a+b就是后一个
		a, b = b, a+b
		return a
	}
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

type intGen func() int

//函数也可以实现接口
func (g intGen) Read(p []byte) (n int, err error) {
	//panic("implement me")
	//调她就取得下一个元素
	next := g()
	//做个限制
	if next > 100000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)
	//把s写进byte哩
	return strings.NewReader(s).Read(p)

}

func main() {
	f := fibonacci()
	printFileContents(f)
}
