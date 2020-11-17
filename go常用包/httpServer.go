package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	//打印请求的主机地址
	fmt.Println(r.Host)

	//打印请求头信息
	fmt.Printf("header content:{%v}\n", r.Header)

	//获取 post请求中 form里边的数据
	fmt.Printf("form content:[%s, %s]\n", r.PostFormValue("username"), r.PostFormValue("passwd"))

	//读取请求信息
	bodyContent, err := ioutil.ReadAll(r.Body)
	if err != nil && err != io.EOF {
		fmt.Printf("read body content failed, err:[%s]\n", err.Error())
		return
	}

	fmt.Printf("body content:[%s]\n", string(bodyContent))
	//返回响应内容
	fmt.Fprintf(w, "hello world!")
}

func main() {
	http.HandleFunc("/index", indexHandler)
	http.ListenAndServe(":8000", nil)
}
