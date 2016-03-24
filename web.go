package main

import (
	"fmt"
	"net/http"
        "os"
)

func sayhello(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()   
	host, err := os.Hostname()
	if err != nil {
            return
        } else {
	    fmt.Fprintf(w, fmt.Sprintf("<html><head></head><body><div align='center'><img height=\"200\" src=\"/static/gaia.jpg\"><h1>Hello world! My hostname is %s </h1></div></body></html>", host))
        }
}


func main() {
        http.HandleFunc("/", sayhello)
	http.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		fmt.Println("Listen And Server", err.Error())
	}
}
