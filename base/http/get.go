package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

func main() {
	// get()
	post()
}

func get() {
	resp, err := http.Get("https://www.jenkins.io/zh/doc/pipeline/tour/running-multiple-steps/")
	if err != nil {
		log.Fatal(err)
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println(string(body))
}

func post() {
	resp, err := http.PostForm("http://www.feijisu4.com/search/", url.Values{"qq": {"golang"}})
	if err != nil {
		log.Fatal(err)
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println(string(body))
}
