package main

import (
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
)

func main() {
	http.HandleFunc("/", commonHandler)
	log.Fatal(http.ListenAndServe("0.0.0.0:80", nil))
}

// common request handler
func commonHandler(writer http.ResponseWriter, request *http.Request) {
	var statusCode = 0
	var responseStr = ""

	if request.URL.Path == "/healthz" {
		statusCode = 200
		responseStr = "the server is ok"
	} else {
		statusCode = 404
		responseStr = "the path is not exist"
	}
	//打印请求ip
	printClientIpAndResponseStatus(request, statusCode)
	//将请求头add到返回头中
	requestHeaderAddResponseHeader(writer, request)
	//将环境变量add到response header中
	addEvnVariableToResponseHeader(writer)
	//设置请求返回码
	writer.WriteHeader(statusCode)
	//设置请求返回参数
	io.WriteString(writer, responseStr)
}

// add environment variable to response header
func addEvnVariableToResponseHeader(writer http.ResponseWriter) {
	os.Setenv("VERSION", "0.0.1")
	version := os.Getenv("VERSION")
	log.Printf("the local ststem var, VERSION: [%s]", version)
	writer.Header().Set("VERSION", version)
}

// print client ip, request path and  response status code
func printClientIpAndResponseStatus(request *http.Request, statusCode int) {
	log.Printf("request client ip :[%s],request path: [%s], response status code : [%d]", getClientIp(request), request.URL.Path, statusCode)
}

// request header add response header
func requestHeaderAddResponseHeader(writer http.ResponseWriter, request *http.Request) {
	for k := range request.Header {
		value := request.Header.Get(k)
		writer.Header().Set(k, value)
	}
}

// GetClientIp /**
func getClientIp(request *http.Request) string {
	ip := request.Header.Get("X-Real-IP")
	if net.ParseIP(ip) != nil {
		return ip
	}
	ip = request.Header.Get("X-Forward-For")
	for _, i := range strings.Split(ip, ",") {
		if net.ParseIP(i) != nil {
			return i
		}
	}
	return strings.Split(request.RemoteAddr, ":")[0]
}
