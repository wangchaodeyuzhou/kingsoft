package main

import (
	"io"
	"log/slog"
	"net/http"
	"time"
)

func main() {
	client := http.Client{Timeout: 5 * time.Second}
	req, err := http.NewRequest("GET", "http://10.228.27.60:7001/timeout", nil)
	if err != nil {
		slog.Error("new request fail", "err", err)
		return
	}
	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		slog.Error("send request fail", "err", err)
		time.Sleep(60 * time.Second)
		return
	}
	defer func() {
		if resp.Body != nil {
			err := resp.Body.Close()
			if err != nil {
				slog.Error("close responseBody body from ai server fail", "err", err)
				return
			}
		}
	}()

	if resp.StatusCode != http.StatusOK {
		slog.Error("send to ai request fail", "http status", resp.StatusCode)
		return
	}
	//	读取响应数据
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		slog.Error("read responseBody body fail", "err", err)
		return
	}
	slog.Info("response", "body", string(body))
	time.Sleep(60 * time.Second)
}
