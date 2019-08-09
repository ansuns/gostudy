package main

import "fmt"

func main() {
	//定义：map[key]val
	//key是无序的
	m := map[string]string{
		"name":    "yangs",
		"course":  "golang",
		"address": "peking",
	}
	fmt.Println(m)

	//第二种方法：使用Make函数
	m2 := make(map[string]int) //空MAP:M2=empty
	fmt.Println(m2)

	//第三种：var
	var m3 map[string]int //M3=NIL
	fmt.Println(m3)

	//遍历
	for k, v := range m {
		fmt.Println(k, v)
	}

	//获取值，如果不存在，拿到的是zero value
	fmt.Println(m["name"])

	//判断key是否存在
	if address, ok := m["address"]; ok {
		fmt.Println(address)
	} else {
		fmt.Println("key no exist")
	}

	//删除
	address, ok := m["address"]
	fmt.Println(address, ok)
	delete(m, "address")
	address, ok = m["address"]
	fmt.Println(address, ok)

}
