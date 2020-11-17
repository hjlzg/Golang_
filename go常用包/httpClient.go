package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	var url = "http://127.0.0.1:8000/index"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("get request failed,err:[%s]", err.Error())
		return
	}

	defer resp.Body.Close()

	bodyContent, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("resp status code:[%d]\n", resp.StatusCode)
	fmt.Printf("resp body data:[%s]\n", string(bodyContent))
	return
}
