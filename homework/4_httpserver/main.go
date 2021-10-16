package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/pprof"
	"os"
	"strings"
)

func main() {
	fmt.Println("starting http server...")

	// 启动一个自定义mux的http服务器
	mux := http.NewServeMux()
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)

	mux.HandleFunc("/", getRootHandler)
	mux.HandleFunc("/healthz", healthz)

	err := http.ListenAndServe(":80", mux)
	if err != nil {
		log.Fatal(err)
	}

}

func getRootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("entering getRequestHeader handler")

	// 1. 接收客户端 request，并将 request 中带的 header 写入 response header
	for k, v := range r.Header {
		//io.WriteString(w, fmt.Sprintf("%s=%s\n", k, v))
		w.Header().Set(k, v[0])
	}

	// 2. 读取当前系统的环境变量中的 VERSION 配置，并写入 response header
	version := os.Getenv("VERSION")
	w.Header().Set("VERSION", version)

	// 3. Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
	returnStatus := http.StatusOK
	w.WriteHeader(returnStatus) // specify the 200 http status code

	IPAddress := r.Header.Get("X-Real-Ip")
	if IPAddress == "" {
		IPAddress = r.Header.Get("X-Forwarded-For")
	}
	if IPAddress == "" {
		IPAddress = r.RemoteAddr
	}
	s := fmt.Sprintf("Client IP is %s\nHttp status code is %d\n", strings.Split(IPAddress, ":")[0], returnStatus)
	fmt.Print(s)
}

// 4. 当访问 localhost/healthz 时，应返回200
func healthz(w http.ResponseWriter, r *http.Request) {
	fmt.Println("entering healthz handler")
	io.WriteString(w, "200\n")
}
