package main

import "fmt"

func length0fNonRepeatingSubStr(s string) int {
	//z-z-zstart-0-1-2-3-4X-y-y-y
	//Oo语言的char是rune类型
	//转成byte：[]byte(s)
	lastOccurred := make(map[rune]int) //int表示每个字母最后出现的位置
	start := 0                         //表示当前找到最长不含有重复字符的字串的开始
	//看看从start到x-1之间有没有x
	maxLength := 0
	for i, ch := range []rune(s) {
		//0下标也是合法，排除掉
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}

		lastOccurred[ch] = i
	}

	return maxLength
}

func main() {
	fmt.Println(length0fNonRepeatingSubStr("ababcabcbbcaabcabcbbabcabcbbbcbb"))
	fmt.Println(length0fNonRepeatingSubStr("啊啊发发发"))
}
