package main

import (
	"golang/errorhandling/fileListeningserver/filelisting"
	"log"
	"net/http"
	"os"
)

type appHandler func(writer http.ResponseWriter, request *http.Request) error

func errWrapper(handler appHandler) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := handler(w, r)
		if err != nil {
			log.Printf("Error occurred handling request: %s", err.Error())
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
func main() {
	http.HandleFunc("/list/",
		//func(writer http.ResponseWriter, request *http.Request) 是一个Handle函数
		errWrapper(filelisting.HandleList))

	err := http.ListenAndServe(":8887", nil)
	if err != nil {
		panic(err)
	}
}
