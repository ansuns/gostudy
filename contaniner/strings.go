package main

import "fmt"

func main() {
	s := "yes我爱我家"
	//rune 转中文
	for i, b := range  []rune(s) {
		fmt.Printf("(%d %c)", i, b)
	}
}
