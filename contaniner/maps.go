package contaniner

import "fmt"

/**
map 是使用哈希表的，key可以比较相等
 */
func main() {
	m := map[string]string {
		"name" : "ansuns",
		"course" : "golang",
		"quality" : "notbad",
	}

	fmt.Println(m)

	//空map:
	m2 := make(map[string]int)
	fmt.Println(m2)

	//空map:
	var m3 map[string]int
	fmt.Println(m3)

	//map遍历
	for k, v := range m {
		fmt.Println(k,v)
	}

	fmt.Println("取值name:")
	fmt.Println(m["name"])

	//删除name
	fmt.Println("delete name:")
	delete(m, "name")
	fmt.Println(m)

}
