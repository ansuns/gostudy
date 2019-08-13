package main

import (
	"fmt"
	"time"
)

/*//channel可以作为参数
func woker(c chan int) {
	for {
		n := <-c //c出来送给n，接收
		fmt.Println(n)
	}
}

func chanDemo() {
	c := make(chan int)
	go woker(c)
	c <- 1 //发一个1
	c <- 2 //发一个2
	time.Sleep(time.Millisecond)
}*/

//chan<- 表示用来收数据
func careteWorker(i int) chan<- int {
	c := make(chan int)
	go worker(i, c)
	return c
}

func chanDemo() {
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		channels[i] = careteWorker(i)
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i;
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i;
	}
}

func worker(i int, c chan int) {
	for {
		fmt.Printf("Worker %d receiver %c\n", i, <-c)
	}
}

//buffer channel
//发数据没有人收是不可以的
func bufferedChannel() {
	c := make(chan int, 3) //3：缓冲区，可以提高性能？
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
}

func main() {
	//chanDemo()
	bufferedChannel()
	time.Sleep(time.Minute)
}
