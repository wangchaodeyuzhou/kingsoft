package main

import (
	"fmt"
	"os"
	"runtime/trace"
)

func main() {
	f, err := os.Create("trace.out")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()

	if err := trace.Start(f); err != nil {
		panic(err)
		return
	}

	defer trace.Stop()

	fmt.Println("hello word ...")
	fmt.Println("go tool trace trace.out")
}
