package mock

import "fmt"

//定义一个结构体
type Retriever struct {
	//内容写在肚子里
	Contens string
}

//实现了Get方法就实现了Stringer接口
func (r *Retriever) String() string {
	return fmt.Sprintf("Retiever:{Conterns=%s}", r.Contens)
}

//实现了Get方法就实现了Poster接口
func (r *Retriever) Post(url string, form map[string]string) string {
	r.Contens = form["name"]
	return "ok"
}

//实现了Get方法就实现了Retriever接口
func (r *Retriever) Get(url string) string {
	return r.Contens
}

//实现一个stinger
