package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func echo(w http.ResponseWriter, r *http.Request) {
	req, _ := ioutil.ReadAll(r.Body)
	fmt.Printf("%s\n", req)

	io.WriteString(w, `{"code": 10, "data": null}`)
}

func main() {
	http.HandleFunc("/client", echo)
	http.HandleFunc("/stream", echo)
	http.HandleFunc("/session", echo)
	http.HandleFunc("/dvr", echo)
	http.HandleFunc("/hls", echo)
	http.ListenAndServe(":8080", nil)
}
