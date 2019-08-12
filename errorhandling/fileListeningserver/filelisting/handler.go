package filelisting

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const prefix = "/list/"

//定义自己的userError的类型
type userError string

func (e userError) Error() string {
	return e.Message()
}

func (e userError) Message() string {
	return string(e)
}

func HandleList(writer http.ResponseWriter, request *http.Request) error {
	//找找看有没有/list/
	if strings.Index(request.URL.Path, prefix) != 0 {
		return userError("Path must be start with " + prefix)
	}
	path := request.URL.Path[len("/list/"):] //list/aa.txt
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	all, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	writer.Write(all)
	return nil
	//把内容写到ResponseWriter哩去

}
