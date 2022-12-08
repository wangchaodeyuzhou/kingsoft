package network

import (
	"fmt"
	"net/http"
)

func myHander(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.RemoteAddr, "连接 success")
	fmt.Println("method, err:", r.Method)

	fmt.Println("url: ", r.URL, "header : ", r.Header, "body : ", r.Body)

	w.Write([]byte("www.baidu,com"))
}
