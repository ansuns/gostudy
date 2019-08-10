package main

import (
	"bufio"
	"fmt"
	"os"
)

func tryDefer() {
	//defer先进后出
	//defer关键字：函数退出时调用
	//也不错怕中间return或出错
	//defer相当于有一个栈，先进后出：3 2 1
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	return
	defer fmt.Println("defer")
}

func writeFile(fileName string) {
	//打开一个文件
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	//只写到了BUFFER里面，要导进去文件
	defer writer.Flush() //defer关键字：函数调用结束写入

	for i := 0; i < 10; i++ {
		fmt.Fprintln(writer, i)
	}
}

func main() {
	tryDefer()
	writeFile("aaa.txt")
}
