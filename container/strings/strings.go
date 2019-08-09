package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "yes我爱你中国!"
	fmt.Println(len(s))
	//X打印字节的具体数字
	fmt.Printf("%X\n", []byte(s))
	for _, b := range []byte(s) {
		fmt.Printf("%X ", b)
	}

	fmt.Println()

	for i, ch := range s {
		fmt.Printf("（%d %X) ", i, ch)
	}

	fmt.Println(utf8.RuneCountInString(s))
	//转换为字节，然后是一个切片slice
	bytes := []byte(s)
	//大于0就可以转为RUNE
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:] //siez的长度已经被转完，所以从size开始
		fmt.Printf("%c ", ch)
	}
	fmt.Println()

	s2 := "yes我爱你中国!"
	//直接转为RUNE
	for i, ch := range []rune(s2) {
		fmt.Printf("(%d %c)", i, ch)
	}
}
