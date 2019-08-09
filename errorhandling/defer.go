package main

import (
	"bufio"
	"fmt"
	"os"
)

func tryDerfer()  {
	defer fmt.Println(1)
	fmt.Println(4)
	defer fmt.Println(2)

	defer fmt.Println(3)
}
func main() {
	tryDerfer()
	writeFile("fib.txt")
}

func fib() func() int {
	a, b := 0, 1
	return func() int{
		a, b = b, a+b
		return a
	}
}

func writeFile(fileName string)  {
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	writer := bufio.NewWriter(file)

	defer writer.Flush()

	//先进后出，写文件
	f := fib()
	for i := 0; i < 90; i++ {
		fmt.Fprintln(writer, f())
	}
}