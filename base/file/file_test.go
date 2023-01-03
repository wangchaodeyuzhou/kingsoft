package file

import (
	"fmt"
	"os"
	"testing"

	"github.com/gogf/gf/v2/text/gstr"
)

func TestFileDirs(t *testing.T) {
	dir, err := os.ReadDir("../channal")
	if err != nil {
		fmt.Println("err : ", err)
		return
	}
	for _, fileInfo := range dir {
		fmt.Println(fileInfo.Name())
	}
}

func TestGstr(t *testing.T) {
	split := gstr.Split("list, common.AttrInfo", ".")
	fmt.Println(split)
	fmt.Println(split[0])
}

type data struct {
	name string
}

func TestMapData(t *testing.T) {
	m := map[string]*data{
		"x": {"JJ"},
	}

	fmt.Println(m["x"])
	m["x"].name = "KK"
	fmt.Println(m["x"])
}
