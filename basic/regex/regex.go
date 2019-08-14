package main

import (
	"fmt"
	"regexp"
)

const text = `my email is abc@qq.com
emainl2 1@qq.com
email3 22@qq.com
`

func main() {
	//re, err := regexp.Compile(text)
	prex := `[a-zA-Z0-9]+@.+\.[a-zA-Z0-9]+` //tab键上面这个
	re := regexp.MustCompile(prex)          //查找abc
	match := re.FindAllString(text, -1)     //在字符串中查找:-1所有
	fmt.Println(match)
}
