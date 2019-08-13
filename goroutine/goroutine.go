package main

import (
	"fmt"
	"time"
)

func main() {
	//goroutine是一种协程(轻量级线程？)
	//main函数也是一个goroutine
	for i := 0; i < 100; i++ {
		//这里，如果直接调用for里面的i，不太安全，所以传参给张哥匿名函数
		go func(i int) {
			//Printf 是一个IO的操作，有一个等待的过程
			fmt.Printf("Hello goroutine:%d\n", i)
			// runtime.Gosched()//非Io操作，交出控制权
		}(i)
	}
	time.Sleep(time.Millisecond)
}
