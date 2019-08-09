package main

import (
	"fmt"
	"golang/retriever/mock"
	"golang/retriever/really"
	"time"
)

const url = "http://www.baidu.com";

type Retriever interface {
	//接口内部函数不需要加func关键字，内部就是函数
	Get(url string) string //有GET方法，接收string，返回sring
}

type Poster interface {
	Post(url string, form map[string]string) string
}

//接口组合
type RetrieverPoster interface {
	Retriever
	Poster
}

func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{"name": "this is anthoer retriver"})
	return s.Get(url)
}

/**
download是使用者
 */
func download(r Retriever) string {
	return r.Get(url)
}

func post(poster Poster) {
	poster.Post(url,
		map[string]string{"name": "yangs", "sex": "man",})
}

func insepct(r Retriever) {
	fmt.Println("Inspecting", r)
	fmt.Printf("%T %v\n", r, r)
	fmt.Print(" > Type switch:")
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents:", v.Contens)
	case *really.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}
	fmt.Println()
}

func main() {
	var r Retriever //
	retriever := mock.Retriever{"this is a retriver"}
	r = &retriever
	insepct(r)
	r = &really.Retriever{
		UserAgent: "Mozilla/5.0",
		Timeout:   time.Minute,
	}
	insepct(r)
	fmt.Println("try a session")
	fmt.Println(session(&retriever))
	//休眠10秒
	time.Sleep(time.Second * 3)
	//fmt.Println(download(r))
	//time.Sleep(time.Second * 5)
}
