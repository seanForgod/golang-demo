package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", handler)
	println("the server is running")
	http.ListenAndServe("localhost:8080", nil)
}

func handler(writer http.ResponseWriter, request *http.Request) {
	log.Printf("http request path: %s", request.URL.Path)
	log.Printf("http request header: %s", request.Header)
	response := ""
	if request.URL.Path == "/test" {
		response = "the server is ok"
	} else {
		response = "the request path is 404.."
	}
	for k := range request.Header {
		writer.Header().Set(k, request.Header.Get(k))
	}
	version := os.Getenv("VERSION")
	log.Printf("the local ststem var, version: %s", version)
	writer.Header().Add("VERSION", version)
	write, err := writer.Write([]byte(response))
	if err != nil {
		log.Println("请求返回异常")
		return
	}
	println(write)

}
