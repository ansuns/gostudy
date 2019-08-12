package main

import (
	"golang/errorhandling/fileListeningserver/filelisting"
	"log"
	"net/http"
	"os"
)

type appHandler func(writer http.ResponseWriter, request *http.Request) error

func errWrapper(handler appHandler) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		//服务器处理函数，也要处理错误，recover
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Paninc:%v", r)
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()
		err := handler(w, r)
		if err != nil {
			log.Printf("Error occurred handling request: %s", err.Error())

			//如果他是一个userError,在这里进行处理
			if userErr, ok := err.(userError); ok {
				http.Error(w,
					userErr.Message(),
					http.StatusBadRequest)
				return
			}
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(
				w,
				http.StatusText(code),
				code)
		}
	}
}

//定义可以给用户看的错误
type userError interface {
	error
	Message() string
}

func main() {
	http.HandleFunc("/",
		//func(writer http.ResponseWriter, request *http.Request) 是一个Handle函数
		errWrapper(filelisting.HandleList))

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
