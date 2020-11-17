package main

import "net/http"

// SayHello is Hello
func SayHello(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello"))
}

func main() {
	http.HandleFunc("/hello", SayHello)
	http.ListenAndServe(":8001", nil)
}
