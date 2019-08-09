package main

import "fmt"

func main() {
	fmt.Println(
		lengthOfNonRepeatingSubStr("abcaa"),
		lengthOfNonRepeatingSubStr("abcdefg"),
		lengthOfNonRepeatingSubStr("aa"),
		lengthOfNonRepeatingSubStr(""),
		lengthOfNonRepeatingSubStr("我爱我家"),
		)
}

func lengthOfNonRepeatingSubStr(s string) int  {
	lastOcur := make(map[rune]int)
	start := 0
	maxLength := 0

	//map 是使用哈希表的，key可以比较相等
	for i, ch := range []rune(s) {

		//判断有没有这个Key
		if lastI, ok := lastOcur[ch]; ok &&  lastI >= start {
			//更新start
			//fmt.Println(i)
			start = lastI + 1
		}

		if i - start + 1 > maxLength {
			maxLength = i - start + 1
		}

		lastOcur[ch] = i
	}

	return maxLength
}
