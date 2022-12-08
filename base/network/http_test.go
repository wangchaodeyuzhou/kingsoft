package network

import (
	"fmt"
	"io"
	"net/http"
	"testing"
)

func TestHttpServer(t *testing.T) {
	http.HandleFunc("/go", myHander)

	http.ListenAndServe(":5000", nil)
}

func TestHttpClient(t *testing.T) {
	resp, _ := http.Get("http://127.0.0.1:5000/go")
	defer resp.Body.Close()

	fmt.Println(resp.Status)
	fmt.Println(resp.Header)

	buf := make([]byte, 1024)
	for {
		n, err := resp.Body.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Println(err)
			return
		} else {
			fmt.Println("read over")
			res := string(buf[:n])
			fmt.Println(res)
			break
		}
	}
}
