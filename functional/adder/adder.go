package adder

import "fmt"

//无参数，返回的是一个函数
func adder() func(int) int {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}

func main() {
	a := adder()
	for i := 0; i < 10; i++ {
		fmt.Println(a(i))
	}
}
