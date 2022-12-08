package json_decoded

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"
)

func TestJson(t *testing.T) {
	uri := "http://ajax.googleapis.com/ajax/services/search/web?v=1.0&rsz=8&q=golang"
	resp, err := http.Get(uri)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}

	defer resp.Body.Close()

	var gr gResponse
	err = json.NewDecoder(resp.Body).Decode(&gr)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}

	fmt.Println(gr)
}

func TestKK(t *testing.T) {
	c := make(map[string]any)
	c["name"] = "Go"
	c["title"] = "programmer"
	c["contact"] = map[string]any{
		"home": "cdsldsdkcsd",
		"cell": "ckdslkcs",
	}

	data, err := json.MarshalIndent(c, "", "	")
	if err != nil {
		return
	}

	fmt.Println(string(data))

	var b bytes.Buffer
	b.Write([]byte("hello"))

	fmt.Fprintln(&b, " world")
	b.WriteTo(os.Stdout)
}

func TestKL(t *testing.T) {
	r, err := http.Get(os.Args[1])
	if err != nil {
		return
	}

	file, err := os.Create(os.Args[2])
	if err != nil {
		return
	}

	defer file.Close()

	dest := io.MultiWriter(os.Stdout, file)
	io.Copy(dest, r.Body)
	if err := r.Body.Close(); err != nil {
		fmt.Println(err)
	}
}
