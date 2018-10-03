package main

import (
	"fmt"
	"project_google_gostudy/retriver/mock"
	"project_google_gostudy/retriver/real"
)

func main() {
	var r Retriver
	r = mock.Retriver{"aaaa"}
	r = real.Retriver{}
	fmt.Println(download(r))
}

type Retriver interface {
	//interface里面全是函数，Get不用加func
	Get(url string) string
}

func download(r Retriver) string {
	return r.Get("http://www.imooc.com");
}
