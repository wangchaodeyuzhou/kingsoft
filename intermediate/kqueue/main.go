package main

import (
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	os.Exit(realMain())
}

func realMain() (exitCode int) {

	return 0
}
