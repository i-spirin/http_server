package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/i-spirin/http_server/config"
)

func hello(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func ip(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "%v\n", req.RemoteAddr)
}

func main() {

	ConfigFilePath := flag.String("c", "config.yaml", "Path to config.yaml")
	flag.Parse()

	conf := config.Config{}
	conf.Parse(*ConfigFilePath)

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/ip", ip)

	http.ListenAndServe(conf.BindHost+":"+fmt.Sprint(conf.BindPort), nil)
}
