package contaniner

import "fmt"

func printSlice(s []int)  {
	fmt.Printf("len s = %d cap s = %d\n", len(s), cap(s))
}
func main() {
	//创建slice
	var s []int //Zero value for slice is nil
	for i := 0; i <= 100; i++ {
		//必须接收append返回值
		if i % 2 == 0 {
			printSlice(s)
			//装不下时，每次扩充 * 2
			s = append(s, i)
		}

	}
	fmt.Println(s)

	//长度是16的slice创建
	s1 := make([]int, 16)
	fmt.Println(s1)

	fmt.Println("copying:")
	//直接覆盖？
	copy(s1, s)
	fmt.Println(s1)

	fmt.Println("deleting:")
	//删除s1的16
	//原理，前面8个 + 加上后面7个 即 s1[:7] + s1[9:],跳掉16
	s1 = append(s1[:8], s1[9:]...)
	fmt.Println(s1)

}
