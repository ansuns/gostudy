package really

import (
	"net/http"
	"net/http/httputil"
	"time"
)

//实现真实的retriever

type Retriever struct {
	//放一些参数
	UserAgent string
	Timeout   time.Duration
}

//用type->Retriever实现Retriever接口
//可以传值也可以传指针
func (r *Retriever) Get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	result, err := httputil.DumpResponse(resp, true)

	//读完要close
	resp.Body.Close()
	if err != nil {
		panic(err)
	}

	return string(result)
}
